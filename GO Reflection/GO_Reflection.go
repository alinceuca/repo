// GO-Reflection
// Link-uri mentionate in prezentare si alte surse de inspiratie:
//
// https://go101.org/article/reflection.html
// https://github.com/a8m/reflect-examples
// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int
	B string
}

type MyInt int

type order struct {
	ordId      int
	customerId int
}

type A struct{}

const n = 255

func main() {

	// --------- Example with numbers ---------

	//var (
	//	a int8
	//	b int16
	//	c uint
	//	d float32
	//	e string
	//)
	//fmt.Println(fill(&a), a)
	//fmt.Println(fill(&b), b)
	//fmt.Println(fill(&c), c)
	//fmt.Println(fill(&d), c)
	//fmt.Println(fill(&e), e)
	//
	//
	//a := 56
	//x := reflect.ValueOf(a).Int()
	//fmt.Printf("type:%T value:%v\n", x, x)
	//b := "Naveen"
	//y := reflect.ValueOf(b).String()
	//fmt.Printf("type:%T value:%v\n", y, y)




	// --------- Example with struct ---------

	//o := order {
	//	ordId: 1234,
	//	customerId: 567,
	//}
	//
	//createQuery(o)
	//createQuery2(o)




	// --------- Example call function ---------

	//// ValueOf returns a new Value, which is the reflection interface to a Go value.
	//v := reflect.ValueOf(A{})
	//m := v.MethodByName("Hello")
	//if m.Kind() != reflect.Func {
	//	return
	//}
	//m.Call(nil)



	// --------- Reflection Law 1 ---------

	//var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x))
	//
	//fmt.Println("value:", reflect.ValueOf(x).String())
	//
	//v := reflect.ValueOf(x)
	//fmt.Println("type: ", v.Type())
	//fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	//fmt.Println("value: ", v.Float())
	//
	//var y uint8 = 'y'
	//w := reflect.ValueOf(y)
	//fmt.Println("type: ", w.Type())                            // uint8.
	//fmt.Println("kind is uint8: ", w.Kind() == reflect.Uint8) // true.
	//y = uint8(w.Uint())
	//
	//var z MyInt = 7
	//q := reflect.ValueOf(z)
	//fmt.Println("type: ", q.Type())
	//fmt.Println("kind: ", q.Kind())
	//// the Kind of q is still reflect.Int, even though the
	//// static type of z is MyInt, not int.
	//// In other words, the Kind cannot discriminate an
	//// int from a MyInt even though the Type can.



	// --------- Reflection Law 2 ---------

	//var x = 3.4
	//v := reflect.ValueOf(x)
	//
	//y := v.Interface().(float64) // y will have type float64.
	//fmt.Println(y)
	//
	//fmt.Println(v.Interface()) 	 //Why not fmt.Println(v)? Because v is a reflect.Value;
	//								// we want the concrete value it holds.



	// --------- Reflection Law 3 ---------

	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	////v.SetFloat(7.1) // Error: will panic.
	//
	//fmt.Println("settability of v:", v.CanSet()) //false
	//
	//p := reflect.ValueOf(&x) // Note: take the address of x.
	//fmt.Println("type of p: ", p.Type())
	//fmt.Println("settability of p: ", p.CanSet())
	////The reflection object p isn't settable, but it's not p we want to set,
	//// it's (in effect) *p. To get to what p points to,
	//// we call the Elem method of Value
	//
	//w := p.Elem()
	//fmt.Println("settability of v:", w.CanSet())
	//
	//w.SetFloat(7.1)
	//fmt.Println(w.Interface())
	//fmt.Println(x)


	// --------- Another example for Law 3 ---------

	//greeting := "hello"
	//f := Foo{A: 10, B: "Salutations"}
	//
	//gVal := reflect.ValueOf(greeting)
	//// not a pointer so all we can do is read it
	//fmt.Println(gVal.Interface())
	//
	//gpVal := reflect.ValueOf(&greeting)
	//// it’s a pointer, so we can change it, and it changes the underlying variable
	//gpVal.Elem().SetString("goodbye")
	//fmt.Println(greeting)
	//
	//fType := reflect.TypeOf(f)
	//fVal := reflect.New(fType)
	//fVal.Elem().Field(0).SetInt(20)
	//fVal.Elem().Field(1).SetString("Greetings")
	//f2 := fVal.Elem().Interface().(Foo)
	//fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
	//





	// --------- reflect.TypeOf() and EXAMINER ---------

	//sl := []int{1, 2, 3}
	//greeting := "hello"
	//greetingPtr := &greeting
	//f := Foo{A: 10, B: "Salutations"}
	//fp := &f
	//
	//slType := reflect.TypeOf(sl)
	//gType := reflect.TypeOf(greeting)
	//grpType := reflect.TypeOf(greetingPtr)
	//fType := reflect.TypeOf(f)
	//fpType := reflect.TypeOf(fp)
	//
	//examiner(slType, 0)
	//examiner(gType, 0)
	//examiner(grpType, 0)
	//examiner(fType, 0)
	//examiner(fpType, 0)




	// --------- Making Without Make ---------

	//// declaring these vars, so I can make a reflect.Type
	//intSlice := make([]int, 0)
	//mapStringInt := make(map[string]int)
	//
	//// here are the reflect.Types
	//sliceType := reflect.TypeOf(intSlice)
	//mapType := reflect.TypeOf(mapStringInt)
	//
	//// and here are the new values that we are making
	//intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	//mapReflect := reflect.MakeMap(mapType)
	//
	//// and here we are using them
	//v := 10
	//rv := reflect.ValueOf(v)
	//intSliceReflect = reflect.Append(intSliceReflect, rv)
	//intSlice2 := intSliceReflect.Interface().([]int)
	//fmt.Println(intSlice2)
	//
	//k := "hello"
	//rk := reflect.ValueOf(k)
	//mapReflect.SetMapIndex(rk, rv)
	//mapStringInt2 := mapReflect.Interface().(map[string]int)
	//fmt.Println(mapStringInt2)

}

func (A) Hello() { fmt.Println("World") }

func DoSomething(f Foo) {
	fmt.Println(f.A, f.B)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

func fill(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", v.Type())
	}
	// get the value that the pointer v points to.
	v = v.Elem()
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.OverflowInt(n) {
			return fmt.Errorf("can't assign value due to %s-overflow", v.Kind())
		}
		v.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.OverflowUint(n) {
			return fmt.Errorf("can't assign value due to %s-overflow", v.Kind())
		}
		v.SetUint(n)
	case reflect.Float32, reflect.Float64:
		if v.OverflowFloat(n) {
			return fmt.Errorf("can't assign value due to %s-overflow", v.Kind())
		}
		v.SetFloat(n)
	default:
		return fmt.Errorf("can't assign value to a non-number type")
	}
	return nil
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
}

func createQuery2(q interface{}){
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}