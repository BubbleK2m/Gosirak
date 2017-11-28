# gosirak

Go 언어 School API 라이브러리 gogyo <https://github.com/DSMdongly> 를 기반으로<br> 제작된 커맨드라인 일일 급식 조회 앱입니다

# usage

gosirak은 다음과 같이 사용할 수 있습니다.<br> gosirak를 이용해 대전의 대덕소프트웨어마이스터고등학교<br>
(구 대덕전자기계고등학교, 학교코드 G100000170)의<br> 10월 30일자 급식을 조회하는 명령어는 다음과 같습니다.

    gosirak -school-kind=high-school -school-region=daejeon -school-code=G100000170 
    -year=2017 -month=10 -day=30

# flags

gosirak의 CommandLine Flags는 다음과 같습니다. 저의 편의를 위해 <br> gosirak은 기본적으로 
대덕소프트웨어마이스터고등학교의 오늘의 식단을 구합니다.

- -school-kind: 학교의 유형 (default: high-school)
- -school-region: 학교의 지역 (default: daejeon)
- -school-code: 학교의 고유 코드 (default: G10000170)
- -year: 구하고자 하는 급식 날짜의 연도 (default: 현재 연도)
- -month: 구하고자 하는 급식 날짜의 월 (default: 현재 월)
- -year: 구하고자 하는 급식 날짜의 일 (default: 현재 일)
