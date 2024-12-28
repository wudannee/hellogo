package ch5

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func DemoStruct() {
	var k []int
	if k == nil {
		fmt.Println("k is nil", k)

		fmt.Println("append 10 to k")
		k = append(k, 10)
		fmt.Println("k is not nil now", k, len(k), cap(k))
	}

	j := make([]int, 0, 5)
	if j != nil {
		fmt.Println("j is not nil", j, len(j), cap(j))
	}

	// create a struct variable
	p := Person{Name: "Susan", Age: 16}
	fmt.Println("p:", p)
	// When a method with a pointer receiver is called on a value, the compiler automatically takes the address of the value and passes a pointer to the method.
	p.SetName("John")
	fmt.Println("john?", p)

	// create a pointer to the struct variable
	fmt.Println("create a pointer to the struct variable")
	pr := &p
	pr.SetName("Lucy")
	fmt.Println("after change p:", pr)

	// var p2 = new(Person)
	// try to call a method on a nil pointer to struct, and fix the panic issue
	var p2 *Person
	operateP2 := func() {
		p2.SetName("Tom")
		fmt.Println("p2:", p2)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
			// possibly p2 is nil, so we need to new it before calling operateP2
			p2 = new(Person)
			operateP2()
		}

		// resume the demo when the panic is handled
		ResumeDemoStruct()
	}()
	operateP2()
}

func ResumeDemoStruct() {
	fmt.Println("resume demo struct")
	playWithTime()

}

func playWithTime() {
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	now = now.In(loc)
	fmt.Println("now:", now)
	name, offset := now.Zone()
	fmt.Println("offset:", offset/3600, name)

	loc = time.FixedZone("", -5*3600)
	now = now.In(loc)
	name, offset = now.Zone()
	fmt.Println("-5,now:", now, "/", name, offset/3600)

	// custom time marshal and unmarshal
	var t CustomTime = CustomTime{Time: now}
	b, err := t.MarshalText()
	if err != nil {
		fmt.Println("time failed to marshal text:", err)
		return
	}
	s := string(b)
	fmt.Println("marshal text:", s)

	t.UnmarshalText([]byte(s))
	fmt.Println("unmarshal text:", t)
	layout :=  "2006.01.02"
	fmt.Printf("and then format it with layout<%s>: %s\n", layout, t.Format(layout))
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalText(b []byte) error {
	t, err := time.Parse("2006/01/02", string(b))
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c *CustomTime) MarshalText() (text []byte, err error) {
	text = []byte(c.Format("2006/01/02"))
	return text, nil
}
