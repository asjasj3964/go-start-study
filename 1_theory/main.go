package main // package와
import (
	"fmt"
	"strings"

	"github.com/asjasj3964/learngo/1_theory/something"
)

// 1-3. function 1
func multiply(a, b int) int { // (a int, b int)와 같다.
	return a * b
}

// 1-3. function 1
// func이 여러 개의 return 값을 가질 수 있다.
func lenAndUpper1(name string) (int, string) { // 문자열의 길이와 대문자 이름 반환
	return len(name), strings.ToUpper(name)
}

// 1-3. function 1
func repeatMe(words ...string) { // 여러 개의 argument를 전달 받는다.
	fmt.Println(words)
}

// 1-4. functions 2
// naked return
func lenAndUpper2(name string) (length int, uppercase string) { // length, uppercase: 반환값
	// defer: func이 끝났을 때 추가적으로 무엇인가 동작하도록 한다.
	defer fmt.Println("lenAndUpper2 is done. ")
	// func이 값을 return 한 후에 실행된다.
	length = len(name) // 다시 생성하는 것이 아닌 업데이트를 하는 것이다.
	uppercase = strings.ToUpper(name)
	return
}

// 1-5. for, range, args
func superAdd(numbers ...int) int {
	total := 0
	// range: array에 loop을 적용할 수 있도록 한다.
	for index := range numbers { // 인덱스를 준다. (0, ..., n - 1)
		// numbers 안에서 조건에 따라 반복 실행을 한다.
		fmt.Println(index)
	}
	for index, number := range numbers { // 인덱스와 값을 반환 받는다.
		fmt.Println(index, number)
	}
	for _, number := range numbers { // 인덱스와 값을 반환 받는다.
		total += number
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

// 1-6. if with a twist
func canIDrink1(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // 조건을 체크하기 전에 변수를 만들 수 있다.
		return false
	} else { // 특이사항: 다음 줄에 쓰면 error
		return true
	}
}

// 1-7. switch
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge { // switch 문을 위한 koreanAge 작성이 가능하다.
	case 10:
		return false
	case 18:
		return true
	}
	return false // 그 외엔 false 반환 (default)
	/* switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false */
}

// 1-11. structs
type person struct {
	name    string
	age     int
	favFood []string
}

func main() { // 그 안의 function을 먼저 찾고 실행시킨다.
	// 1-0. main package
	fmt.Println("Hello World!")
	fmt.Println("--------------------")

	// 1-1. packages and imports
	something.SayHello() // 대문자로 시작 -> 패키지로부터 exprot된 func
	// something.sayBye() // 소문자로 시작 -> private이므로 실행시킬 수 없다.
	fmt.Println("--------------------")

	// 1-2. variables and constrants
	const name1 string = "seongjin" // 상수
	// name = "heejin" // 에러, 값을 바꿀 수 없다.
	fmt.Println(name1)
	name2 := "seongjin" // type은 임의로 변경이 가능하다(첫번째 값 기준). 축약형은 func 안에서만 가능하고 변수에만 적용 가능하다.
	// var name2 string = "seongjin" // 변수
	name2 = "heejin" // 값을 바꿀 수 있다.
	fmt.Println(name2)
	fmt.Println("--------------------")

	// 1-3. function 1
	fmt.Println(multiply(3, 4))
	totalLength, upperName := lenAndUpper1("seongjin") // 정의하고 사용하지 않으면 에러를 발생시킨다. 둘 중 하나만 반환 받을 수 없고 둘 다 반환 받아야 한다.
	fmt.Println(totalLength, upperName)
	totalLength, _ = lenAndUpper1("seongjin") // _: 무시된 value
	fmt.Println(totalLength)
	repeatMe("seongjin", "beomgeun", "edword", "bella")
	fmt.Println("--------------------")

	// 1-4. functions 2
	totalLength, upperName = lenAndUpper2("seongjin")
	fmt.Println(totalLength, upperName)
	fmt.Println("--------------------")

	// 1-5. for, range, args
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)
	fmt.Println("--------------------")

	// 1-6. if with a twist
	fmt.Println(canIDrink1(19))
	fmt.Println("--------------------")

	// 1-7. switch
	fmt.Println(canIDrink2(18))
	fmt.Println("--------------------")

	// 1-8. pointers
	a := 2
	b := &a // &: address, a를 살펴보는 pointer. a의 메모리 주소에 접근한다.
	fmt.Println(&a, b)
	fmt.Println(*b) // *: 메모리 주소를 살펴본다.
	// 값을 복사시키는 것이 아닌 메모리에 저장된 object를 서로 똑같이 가지고 싶도록 한다.
	a = 5
	fmt.Println(*b) // a의 메모리를 살펴보는 것이므로 값이 바뀐다.
	*b = 20         // b가 a의 주소와 연결되어 있으므로 a의 값도 바뀐다.
	fmt.Println("--------------------")

	// 1-9. arrays and slices
	// array - 배열의 크기를 제한할 때
	names1 := [5]string{"seongjin", "beomgeun", "edword", "bella"}
	names1[3] = "nico"
	fmt.Println(names1)
	names1[4] = "heejin"
	fmt.Println(names1)
	// slices - 배열의 크기에 제한 없이 요소를 추가하고 싶을 때
	// length가 없다.
	names2 := []string{"seongjin", "beomgeun", "edword", "bella"}
	names2[2] = "abdulah"
	fmt.Println(names2)
	// names2[4] = "heejin" // slice는 이런 식으로 추가하지 않는다.
	names2 = append(names2, "heejin")
	fmt.Println(names2)
	fmt.Println("--------------------")

	// 1-10. maps
	students := map[string]string{"name": "seongjin", "age": "25"} // map[key]value{key1:value1, key2:value2, ...}
	fmt.Println(students)                                          // kwy를 기준으로 오름차순 정렬된다.
	for key, value := range students {
		fmt.Println(key, value)
	}
	fmt.Println("--------------------")

	// 1-11. structs
	// map에 원소 추가
	// 어떤 원소가 map에 존재하고 있는지 확인할 때
	favFood := []string{"Gambas", "Hamburger"}
	beomgeun := person{"beomgeun", 25, favFood} // 작성 방법 1
	fmt.Println(beomgeun)
	fmt.Println(beomgeun.age)
	fmt.Println(beomgeun.favFood)
	beomgeun = person{name: "beomgeun", age: 26, favFood: favFood} // 작성 방법 2
	fmt.Println(beomgeun)
	fmt.Println("--------------------")
}
