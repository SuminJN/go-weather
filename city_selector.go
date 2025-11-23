package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cityOption struct {
	nameKR string
	nameEN string
}

var cityOptions = []cityOption{
	{nameKR: "서울", nameEN: "Seoul"},
	{nameKR: "부산", nameEN: "Busan"},
	{nameKR: "인천", nameEN: "Incheon"},
	{nameKR: "대구", nameEN: "Daegu"},
	{nameKR: "대전", nameEN: "Daejeon"},
	{nameKR: "광주", nameEN: "Gwangju"},
	{nameKR: "울산", nameEN: "Ulsan"},
	{nameKR: "제주", nameEN: "Jeju"},
}

var cityLookup = buildCityLookup()

func buildCityLookup() map[string]cityOption {
	lookup := make(map[string]cityOption, len(cityOptions)*3)
	for _, option := range cityOptions {
		lookup[option.nameKR] = option
		lookup[option.nameEN] = option
		lookup[strings.ToLower(option.nameEN)] = option
	}
	return lookup
}

func determineCity(args []string) (string, error) {
	if len(args) > 1 {
		return resolveCityArg(args[1]), nil
	}
	return promptCitySelection()
}

func resolveCityArg(input string) string {
	trimmed := strings.TrimSpace(input)

	if option, ok := cityLookup[trimmed]; ok {
		fmt.Printf("입력받은 도시: %s (%s)\n", option.nameKR, option.nameEN)
		return option.nameEN
	}

	lower := strings.ToLower(trimmed)
	if option, ok := cityLookup[lower]; ok {
		fmt.Printf("입력받은 도시: %s (%s)\n", trimmed, option.nameEN)
		return option.nameEN
	}

	fmt.Printf("입력받은 도시: %s\n", trimmed)
	return trimmed
}

func promptCitySelection() (string, error) {
	fmt.Println("조회 가능한 도시 예시:")
	for i, option := range cityOptions {
		fmt.Printf("  %d. %s (%s)\n", i+1, option.nameKR, option.nameEN)
	}
	fmt.Printf("도시 번호를 입력하세요 (Enter 시 %s 선택): ", cityOptions[0].nameKR)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("입력 읽기 오류: %w", err)
	}

	choice := strings.TrimSpace(input)
	if choice == "" {
		selected := cityOptions[0]
		return selected.nameEN, nil
	}

	idx, err := strconv.Atoi(choice)
	if err != nil || idx < 1 || idx > len(cityOptions) {
		return "", fmt.Errorf("유효한 번호를 입력해주세요")
	}

	selected := cityOptions[idx-1]
	return selected.nameEN, nil
}
