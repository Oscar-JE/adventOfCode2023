package parttwo

// func TestCaseInstruktion1(t *testing.T) {
// 	t.Skip()
// 	f := field.Init("???.###")
// 	o := order.Init([]int{1, 1, 3})
// 	res := NumberOfCombination(f, o)
// 	if res != 1 {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara 1", res)
// 	}
// }

// func TestCaseInstruktion2(t *testing.T) {
// 	f := field.Init(".??..??...?##.")
// 	o := order.Init([]int{1, 1, 3})
// 	res := NumberOfCombination(f, o)
// 	if res != 16384 {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, 16384)
// 	}
// }

// func TestCaseInstruktion3(t *testing.T) {
// 	f := field.Init("?#?#?#?#?#?#?#?")
// 	o := order.Init([]int{1, 3, 1, 6})
// 	res := NumberOfCombination(f, o)
// 	expected := 1
// 	if res != expected {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, expected)
// 	}
// }

// func TestCaseInstruktion4(t *testing.T) {
// 	f := field.Init("????.#...#...")
// 	o := order.Init([]int{4, 1, 1})
// 	res := NumberOfCombination(f, o)
// 	expected := 16
// 	if res != expected {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, expected)
// 	}
// }

// func TestCaseInstruktion5(t *testing.T) {
// 	f := field.Init("????.######..#####.")
// 	o := order.Init([]int{1, 6, 5})
// 	res := NumberOfCombination(f, o)
// 	expected := 2500
// 	if res != expected {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, expected)
// 	}
// }

// func TestCaseInstruktion6(t *testing.T) {
// 	f := field.Init("?###????????")
// 	o := order.Init([]int{3, 2, 1})
// 	res := NumberOfCombination(f, o)
// 	expected := 506250
// 	if res != expected {
// 		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, expected)
// 	}
// }

// func TestInputLine4(t *testing.T) {
// 	t.Skip("takes to long time")
// 	f := field.Init(".#????.????#??????")
// 	o := order.Init([]int{2, 1, 1, 2, 1, 1})
// 	res := NumberOfCombination(f, o)
// 	fmt.Println(res)
// }
