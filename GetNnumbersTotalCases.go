package mathlogic

// n개의 숫자의 합이 total을 이루는 모든 경우를 구한다 (n > 0, total >= 0), {1,2}와 {2,1}은 서로 다른 경우이다.
func GetNnumbersTotalCases(n, total int) [][]int {

	// case : {total , 나머지 case[1:]에서 n-1개의 숫자의 합이 0을 이루는 경우1}, {total, ~경우2}, {total, ~경우3}, ...
	//        {total - 1 , 나머지 case[1:]에서 n-1개의 숫자의 합이 1을 이루는 경우1}, {total-1, ~경우2}, {total-1, ~경우3}, ...
	//        {total - 2 , 나머지 case[1:]에서 n-1개의 숫자의 합이 2를 이루는 경우1}, {total-2, ~경우2}, {total-2, ~ 경우3}, ...
	//            ...
	//        {0, 나머지 case[1:]에서 n-1개의 숫자의 합이 total을 이루는 경우1}, {0, ~경우2}, {0, ~경우3}, ...

	result := [][]int{}
	aCase := []int{}
	if n == 1 { // 1개 숫자의 합이 total인 경우는 한 경우 밖에 없음. 그 경우를 리턴.
		aCase = append(aCase, total)
		return [][]int{aCase}
	}
	for i := 0; i <= total; i++ {
		bCase := GetNnumbersTotalCases(n-1, i) // n-1개의 숫자의 합이 i를 이루는 모든 경우가 리턴되어 저장됨
		for _, value := range bCase {
			aCase = []int{}
			aCase = append(aCase, total-i)  // n개의 숫자 중 첫번째 숫자는 total - i,
			aCase = append(aCase, value...) // 두번째 이후로 {total - i인 첫번째 숫자, n-1개의 숫자의 합이 i를 이루는 한 경우}...
			result = append(result, aCase)  // n개의 숫자의 합이 total을 이루는 한 경우를 저장
		} // 반복문을 돌려서 첫번째 숫자가 total-i 인 모든 경우를 result에 저장
	} // 모든 i값에 대하여 반복하며 n개 숫자의 합이 total 인 모든 경우를 result에 저장
	return result // 저장된 모든 경우를 리턴
}
