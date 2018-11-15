# Go (Golang) Json serialization and deserialization

This is an example of serialization and deserialization of JSON using Go (Golang). It shows how to best manage null elements, absent elements, recursive structures. During one of the last hack I had to participate we had to handle large, recursive JSON which included null elements and elements which needed to be serialized as null but as well others which needed to be fully absent.

While Golang support Json deserialization and serialization, you'll need to "Marshal" every type to map them on a system type. The system types in go have the notion of null, it's *nil*, but it doesn't support an Int or a String to be *nil*.

## Getting Started

All this example has been build using the excellent [VS Code](https://code.visualstudio.com/Download) and the Go extension. It does support perfectly running and debugging Go projects on Windows, Linux and Mac.

Note that all the code is on the *[main.go](main.go)* file, you can directly run it and follow the results in this explanation.

## The basic deserialization/serialization in GO

We will use a simple Employee class that will have few properties and being recursive.

```go
// Employee3 class example
type Employee3 struct {
    Name     string    `json:"name"`
    Age      int       `json:"age"`
    Salary   int       `json:"salary"`
    Employee *Employee3 `json:"employee"`
    Other    *Other    `json:"other"`
}

// Other structure
type Other struct {
    Tagada string `json:"tagada,omitempty"`
}
```

In Go, all public properties or methods has to start with a upper case character, all internal ones with a lower case. So in those 2 simple classes, all elements are public.

You will note that both the *Employee* and *Other* properties from the main *Employee3* structure are using a star, which means they are using pointes. It is not mandatory for the *Other* one but it is mandatory for the *Employee* one as Go does not support recursivity.

Note that you'll have to specify `json:"name"` to map the json element with a property. This allow to have a different name for example.

### Let's first serialize the json

```go
    empobj3 := Employee3{Name: "", Age: 24, Salary: 344444}
    emp3, _ := json.Marshal(empobj3)
    fmt.Println(string(emp3))
```

As a result we will get:

```json
{"name":"","age":24,"salary":344444,"employee":null,"other":null}
```

So with no surprise, the serialization looks correct and *employee* ad *other* are null. Let's wait to see how to not serialize them if needed later on.

### Let's now deserialize the json

```go
    dt := "{\"name\": null, \"age\": null, \"salary\": 1234, \"employee\": { \"name\": \"tagada\", \"age\": 25, \"salary\": 34567}}"
    emptd := []byte(dt)
    fmt.Println(string(emptd))
    json.Unmarshal(emptd, &retobj3)
    fmt.Println("From Employee3, Age, Name, Employee and Other")
    fmt.Println(retobj3.Age)
    fmt.Println(retobj3.Name)
    fmt.Println(retobj3.Employee)
    fmt.Println(retobj3.Other)
```

And as a result, we will see:

```
From Employee3, Age, Name, Employee and Other
0
&{tagada 25 34567 <nil> <nil>}
<nil>
```

So now let's reserialize it:

```go
    ser3, _ := json.Marshal(retobj3)
    fmt.Println(string(ser3))
```

And check the result:

```json
{"name":"","age":0,"salary":1234,"employee":{"name":"tagada","age":25,"salary":34567,"employee":null,"other":null},"other":null}
```

From the original serialized elements, we can see that it looks good except that the *other* and sub *employee* elements are now showing as null.

So how to better manage *null* or empty values in Go for JSON deserialization and serialization?

## Behind the scene

Go is using *Marshal* and *Unmarshal* to serialize and deserialize every element. Later on, I'll use the package *gopkg.in/guregu/null.v3* so we can use it as an example here to understand how it's working.

The package implement all the JSON serialization and deserialization marshaling functions. Go has an explicit notion of interface, if the definition is the same as in one of the package it import, then, the function implemented will be used.

What does this package do is creating specific packages *null* and *zero* which are implementing types that can support null and have a specific behavior when getting serialized into/deserialized from json with the marshalling functions.

in short, we will have a type *null.String*, *null.Int* which we will use a bit later and play with. Those are especially built when you need to connect to SQL databases and you'll need this nothing of null and you'll need to serialize as null or zero.

## Let's serialize various structures with diverse options

I've created multiple structures so we can mix the tests and see the results.

```go
package main

import (
    "encoding/json"
    "fmt"

    null "gopkg.in/guregu/null.v3"
)

// Employee class example
type Employee struct {
    Name   null.String `json:"name,omitempty"`
    Age    null.Int    `json:"age"`
    Salary int         `json:"salary"`
    // In GO, no recursivity is permitted,
    // So need to use pointers to avoid it
    Employee *Employee `json:"employee"`
    Other    *Other    `json:"other"`
}

// Employee2 class example
type Employee2 struct {
    Name     string     `json:"name,omitempty"`
    Age      int        `json:"age,omitempty"`
    Salary   int        `json:"salary,omitempty"`
    Employee *Employee2 `json:"employee,omitempty"`
    Other    *Other     `json:"other,omitempty"`
}

// Employee3 class example
type Employee3 struct {
    Name     string     `json:"name"`
    Age      int        `json:"age"`
    Salary   int        `json:"salary"`
    Employee *Employee3 `json:"employee"`
    Other    *Other     `json:"other"`
}

// Employee4 class example
type Employee4 struct {
    Name     string     `json:"name,omitempty"`
    Age      int        `json:"age,omitempty"`
    Salary   int        `json:"salary,omitempty"`
    Employee *Employee4 `json:"employee,omitempty"`
    Other    Other      `json:"other,omitempty"`
}

// Other structure
type Other struct {
    Tagada string `json:"tagada,omitempty"`
}
```

For the serialization and deserialization, you can use the attribute *omitempty* which will not serialize an element if it's null or empty. I've been mixing the specific *null* types as well as normal ones, mixing as well the *Other* class with pointers and without, with and without the *omitempty* attribute.

The rest of the code, just create various objects and serialize them
```go
empobj := Employee{Age: null.IntFrom(24), Salary: 344444}
    empobj := Employee{Age: null.IntFrom(24), Salary: 0}
    empobj2 := Employee2{Name: "", Age: 24, Salary: 0}
    empobj3 := Employee3{Name: "", Age: 24, Salary: 0}
    empobj4 := Employee4{Name: "", Age: 24, Salary: 0}
    emp, _ := json.Marshal(empobj)
    emp2, _ := json.Marshal(empobj2)
    emp3, _ := json.Marshal(empobj3)
    emp4, _ := json.Marshal(empobj4)
    fmt.Println(string(emp))
    fmt.Println(string(emp2))
    fmt.Println(string(emp3))
    fmt.Println(string(emp4))
```

when running this code, we'll have a as result:

```json
{"name":null,"age":24,"salary":0,"employee":null,"other":null}
{"age":24}
{"name":"","age":24,"salary":0,"employee":null,"other":null}
{"age":24,"other":{}}
```

Couple of observations for the *Employee* struct:

- we're using the *null.String* type and as expected the result of the empty string is a null once serialized.
- ho, wait! but we've specified the *omitempty* attribute, so what's happening? Well, looks like it's a bug from Go, see [open issue](https://github.com/golang/go/issues/11939)
- so we can't use *null.String* to carry a null or empty and not serialize it. We'll see later on a solution for that
- as expected, the *employee* and *other* elements are serialized as null as both are *nil*

Couple of observations for the *Employee2* struct:

- in this case, we only serialize what is not null or not empty, in our case we then just have *age*
- ho wait! but I wangted the salary to be serialized as 0! Well, read the first part, if you want it to be serialized as 0, don't use the *omitempty* attribute
- great but what if it comes from a deserialization, where the element was null, it will be deserialized as 0 and if not using the *omitempty*, it will then be serialized as 0 which is not what was wanted! So in this case, use the *null.Int* and it will always be serialized.

Couple of observations for the *Employee3* struct:

- in this one, all is serialized whatever happen
- string will be as empty quotes, pointers as null

Couple of observations for the *Employee4* struct:

- this one is close to the second case as we've been using the *omitempty* on all the elements. The only difference is the *Other* struct which is not a pointer. And in this case, it will be serialized whatever happens. 
- but what's in the struct as it's an empty string won't be serialized because we've been using the *omitempty* attribute on the *Tagada* property
- so question: how then would this serialization be if we were not using the *mitempty* attribute? Easy, it will serialize it: ```{"age":24,"other":{"tagada":""}}``` because it's not a pointer, the object is created in memory even if it doesn't exist.

### Summary for the serialization

- use the *omitempty* attribute not to serialize text which are blank or empty
- be careful when using the *omitempty* attribute with numbers, a 0 value will not be serialized
- use the structures like *null.String* or *null.Int* or *zero.Int* when you have to serialize even if empty or null or 0.
- the *omitempty* is not working on structures. Only on pointers! So play with that in your structures to force one or the other behavior

## Let's deserialize various structures

Now, we will use a test text for the deserialization of our Json and see the result on the various classes. Plus we will serialize it after and compare with the original.

The basic code is the same for all the 4 Employees structures. It's just running with the names changed for the 4 structures:

```go
dt := "{\"name\": null, \"age\": null, \"salary\": 1234, \"employee\": { \"name\": \"tagada\", \"age\": 25, \"salary\": 34567}}"
emptd := []byte(dt)
fmt.Println(string(emptd))
var retobj Employee
json.Unmarshal(emptd, &retobj)
fmt.Println("From Employee, Age, Name, Employee and Other")
fmt.Println(retobj.Age)
fmt.Println(retobj.Name)
fmt.Println(retobj.Employee)
fmt.Println(retobj.Other)
```

Results are:

```
From Employee, Age, Name, Employee and Other
{{0 false}}
{{ false}}
&{{{tagada true}} {{25 true}} 34567 <nil> <nil>}
<nil>
From Employee2, Age, Name, Employee and Other
0

&{tagada 25 34567 <nil> <nil>}
<nil>
From Employee3, Age, Name, Employee and Other
0

&{tagada 25 34567 <nil> <nil>}
<nil>
From Employee4, Age, Name, Employee and Other
0

&{tagada 25 34567 <nil> {}}
{}
```

Couple of comments on the *Employee* structure:

- we're using the *null.String* and *null.Int* structure. When printing them, their representation is the elements that are in the class. so the age is initialized at 0 and has no value. The string is initialized to an empty string and has no value.
- as the same principle the sub *employee* element has a sub element *other* and the value for the *name* is *tagada* and has a value, the *age* is 25 and has a value as well
- those *null* and *zero* structures clearly allow to keep track of what's happening, changing the 0 for something will directly change the false value to true as we'll have put a value. 

Couple of comments on the *Employee2* and *Employee3* structures:

- The *age* is 0 as it was null, the *name* is empty as it was null, both sub elements are *nil* as they were null
- the *employee* sub element is as well as expected
- for the subelement, note the *&* as it's a pointer. Go is nice enough to display the value of the element it is pointing on in case of a pointer

Couple of comments on the *Employee4* structure:

- Everything is as expected, because the *other* element is not a pointer, it has been initialized and does appear as an empty element. But reality is that its only string element is an empty string initialized as well.

So for deserialization, there is nothing surprising, all as expected. Bottom line, except is you need to take track of an empty or null element, not need to use the *null.String* and other structures like this. But as soon as you'll need, use them!

And what if we serialize all those elements now and compare with the initial string?

To serialize those elements, we'll just run it 4 times with the correct variable names:

```go
    fmt.Println("From Employee, reserializing it")
    ser, _ := json.Marshal(retobj)
    fmt.Println(string(ser))
```

As a result, we'll have the 4 following json:

```json
"name":null,"age":null,"salary":0,"employee":{"name":"tagada","age":25,"salary":34567,"employee":null,"other":null},"other":null}
{"employee":{"name":"tagada","age":25,"salary":34567}}
{"name":"","age":0,"salary":0,"employee":{"name":"tagada","age":25,"salary":34567,"employee":null,"other":null},"other":null}
{"employee":{"name":"tagada","age":25,"salary":34567,"other":{}},"other":{}}
```

And without any surprise, the behavior is exactly what was expected and explained in the previous part.

### Summary of the deserialization

- if you don't need to reserialize after and if you don't need to take track of the initial if it was null, then you can use the normal types and you don't need to use the pointers for the sub objects (you still need if they are recursive)
- if you need to serialize after or keep track of a null, then use a mix of structures like *null.Int* and pointers for sub objects or sub elements

## Conclusion

Serializing and deserializing with Go and keeping track of the initial value to reserialize a null as a null is not that trivial. Especially if some of the elements do not need to be serialized.

It's a case by case approach, mixing pointers, *omitempty* attributes and usage of structures like *null.String* or *zeor.String*.

