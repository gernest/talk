# Improving gorm

For the impatient. 
[ here is the source code for the improved gorm](https://github.com/gernest/ngorm)

__TL:DR__

One redditor in critique of my web framework mentioned that using gorm as the main ORM was bad, the reason he said gorm was bloated. Since I don't beleive in heresay, I decided to find out for myself.

Before I go on and highlight the things that I found that I din't like and thought they could get some improvement, I would like to outline things that I like about gorm( In its current state).

- Clean API : The gorm API is elegant indeed, you get almost all the basic things you might need to interact with a relational database. Take for instance `db.First(&user)` to find the first user. `db.Automigrate(&User{})` to automaticallly migrate user model. I don't regret picking this ORM I will still pick it when given a choice 1000 times.

- Intuitive mapping to Go types: Mapping of results just works. You suplly a pointer to your models and it is magically filled with results.

- It just works: This is one of the libraries that I don't have to worry about, It works like a charm.

Okay, those are some of the coll stuffs I like about gorm.

# The parts I found that could use some improvement

### Heavy reliance on globals that can lead to unpredictable behaviur

Take for instance the function for updating time https://github.com/jinzhu/gorm/blob/0fbff1e8f0821fc67ef793a9f403beda5ca372b8/utils.go#L21
This function is used everywhere in gorm and it is exported , so any part of your application can change this and you will have inconstent time value. This can be seen as a feeature, but I don't think it is a good idea, and here is why.

```go
func CrazyTime() time.Time {
	src := "Mon Jan 2 15:04:05 -0700 MST 2006"
	t, _ := time.Parse(time.ANSIC, src)
	return t
}
gorm.NowFunc=CrazyTime
```

Adding this function anywhere in your application will result into confusing time values depending where you called it, how and when you did that.

For isntance https://github.com/jinzhu/gorm/blob/9edd66250e8ae11d572213054643b7bb1ce4d102/callback_create.go#L33-L37

```go
func updateTimeStampForCreateCallback(scope *Scope) {
	if !scope.HasError() {
		now := NowFunc()
		scope.SetColumn("CreatedAt", now)
		scope.SetColumn("UpdatedAt", now)
	}
}
```
That is where CreatedAT and UpdateAt time stamps are added. Guess what, with our `CrazyTime` function both fields will be not be accurate, well they will all contain the value returned by `CrazyTime`. This happens all over the code base multiple places relies on this.

The reliance on global scope doesn't end here. I will need another post, maybe even longer to highlight this.

## Project structure

This is true for many Go projects which have been around for a long time. Having everything under one roof with this kind of codebase makes it hard to navigate, and result in obsecure naming( you can possibly out of names! lol!). `scope`,`search` and `DB` are all mixed together it is hard to follow when trying to make changes. And difficult to know what happens where and when.

## Execution model

Simply there is no execution model. Anything can happen anywhere and god knows how! I wont go into details abouut how callbacks come into the picture.

Okay, not to bore you I did try to improve this. Address all the issues. The changes were breaking the API so I didn't bother the author first, I just forked and started to work on it.

These are highlights on how some of the issues I highlighted here were addressed. 

__NOTE__ : I implemented a dialtect for ql database, since I needed something fast and easy to experiment with my ideas ql was the best thing I can find, the goal was not to support everything lets stick with that first.

## Eliminating global scope
Yep, globals are gone, the problem of unpredictability is gone too. Instead anything that can be shared is passed explicityl through the `engine struct`

```go
//Engine is the driving force for ngorm. It contains, Scope, Search and other
//utility properties for easily building complex SQL queries.
type Engine struct {
	RowsAffected int64

	//When this field is set to true. The table names will not be pluarized.
	//The default behaviour is to plurarize table names e.g Order struct will
	//give orders table name.
	SingularTable bool
	Ctx           context.Context
	Dialect       dialects.Dialect

	Search    *model.Search
	Scope     *model.Scope
	StructMap *model.SafeStructsMap
	SQLDB     model.SQLCommon
	Log       *logger.Zapper

	Now func() time.Time
}
```

See the now func? Yeah, if you go on and change it, only the functions that will be applied with this engine will be affected. There is more into this story, and It is still a work in progress, so a lot of room for improvement

## Fix project structure
```
.
├── builder
│   ├── sql.go
│   └── sql_test.go
├── dialects
│   ├── ql
│   │   ├── ql.go
│   │   └── ql_test.go
│   └── dialect.go
├── engine
│   ├── db.go
│   └── db_test.go
├── errmsg
│   └── errors.go
├── fixture
│   └── fixtures.go
├── hooks
│   ├── default.go
│   └── hooks.go
├── logger
│   └── logger.go
├── model
│   ├── field.go
│   ├── field_test.go
│   └── struct.go
├── regexes
│   └── regexes.go
├── scope
│   ├── join_table.go
│   ├── scope.go
│   └── scope_test.go
├── search
│   ├── search.go
│   └── search_test.go
├── util
│   ├── utils.go
│   └── utils_test.go
├── GUIDE.md
├── License
├── README.md
├── examples_test.go
├── ngorm.go
├── ngorm_test.go
└── txt

13 directories, 30 files
```

## Fix execution model

There are two phases, SQL generation and SQL execution. The areas which builds SQL are in `builder` subpackcge. Execution is handled in expicit manner, in way that you can easy know who is executing what and when.


__SUMMARY__: gorm needs improvement not bashing without any evidence!

__AUTHOR_NOTE__: It was fun having an excuse to play with ql database. I would like to see the efforts to bechmark my implementation so as to make it easy to add performance improvement since every function is known, and well documented and there is no any magical tricts happening.

Unfortunate, I can't keep hacking on this. I have to go back to learning algorithms, it saddens me that startups have a fetish for algorithms it is okay though. It was a fun endeavor!

[source code to the gorm with improvement](https://github.com/gernest/ngorm)
Enjoy!
