package tests

import (
	"net/http"
)

type TestTableStruct = struct {
	name         string
	requestBody  string
	expectedCode int
}

// Тесты на /get_token
func getTestTableCreateNewSong() []TestTableStruct {
	testTableGetToken := []TestTableStruct{
		{
			name: "Test #1. OK. Корректный стандартный запрос",
			requestBody: `{
			"group":"test",
			"song":"test",
			"link":"test",
			"text":"test",
			"release_date":"26.08.2005"
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #2. NOT OK. Неверное release_date",
			requestBody: `{
			"group":"test",
			"song":"test1",
			"link":"test",
			"text":"test",
			"release_date":"фффф"
			}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Test #3. NOT OK. Неверное request body",
			requestBody: `{
				"ababab":"babab"
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #4. NOT OK. пустые поля",
			requestBody: `{
			"group":"",
			"song":"",
			"link":"",
			"text":"",
			"release_date":"26.08.2005"
			}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	return testTableGetToken
}

// Тесты на /get_all_songs
func getTestTableGetAllSongs() []TestTableStruct {
	testTableGetToken := []TestTableStruct{
		{
			name: "Test #1. OK. Корректный стандартный запрос",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"",
			"field":"",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #2. NOT OK. Неверное direction (page 1, direction previous)",
			requestBody: `{
			"direction":"previous",
			"page":1,
			"filter":"",
			"field":"",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #3. OK. Запрос на существующую следующую страницу (page 1, direction next)",
			requestBody: `{
			"direction":"next",
			"page":1,
			"filter":"",
			"field":"",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #4. NOT OK. пустые поля",
			requestBody: `{
			"direction":"",
			"page":0,
			"filter":"",
			"field":"",
			"count_songs_on_page":0
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #5. NOT OK. несуществующая страница",
			requestBody: `{
			"direction":"",
			"page":99999999999,
			"filter":"",
			"field":"",
			"count_songs_on_page":0
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #6. NOT OK. некоректный json (не все поля)",
			requestBody: `{
			"direction":"",
			"page":99999999999
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #7. NOT OK. некоректный json (есть ненужные поля)",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"",
			"field":"",
			"count_songs_on_page":0,
			"a":"a"
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #8. OK. Корректный стандартный запрос c фильтром по имени",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"asc",
			"field":"name",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #9. OK. Корректный стандартный запрос c фильтром по группе",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"asc",
			"field":"group",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #10. Not OK. Не корректный стандартный запрос c фильтром по неизвестному параметру",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"asc",
			"field":"groupa",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #11. Not OK. Не корректный стандартный запрос c фильтром по сортировка с неизвестным параметром",
			requestBody: `{
			"direction":"",
			"page":1,
			"filter":"ascanddesc",
			"field":"group",
			"count_songs_on_page":2
			}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	return testTableGetToken
}

// Тесты на /info
func getTestTableInfo() []TestTableStruct {
	testTableGetToken := []TestTableStruct{
		{
			name: "Test #1. OK. Корректный стандартный запрос",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities"
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #2. NOT OK. Несуществующая group",
			requestBody: `{
			"group":"basdadlasld",
			"song":"Dark Necessities"
			}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Test #3. NOT OK. Несуществующая song",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"asdasdasd"
			}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Test #4. NOT OK. пустые поля",
			requestBody: `{
			"group":"",
			"song":""
			}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	return testTableGetToken
}

// Тесты на /delete_song
// Удаление подразумевает что есть песня и группа со значениями test
func getTestTableDeleteSong() []TestTableStruct {
	testTableGetToken := []TestTableStruct{
		{
			name: "Test #1. OK. Корректный стандартный запрос",
			requestBody: `{
			"group":"test",
			"song":"test"
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #2. NOT OK. Неверное group",
			requestBody: `{
			"group":"aaaaaaaaaaaaaaaa",
			"song":"test"
			}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Test #3. NOT OK. Неверное song",
			requestBody: `{
			"group":"test",
			"song":"aaaaaaaaaaaaaaaaaaaa"
			}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Test #4. NOT OK. пустые поля",
			requestBody: `{
			"group":"",
			"song":""
			}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	return testTableGetToken
}

// Тесты на /get_all_songs
func getTestTableSongByVerses() []TestTableStruct {
	testTableGetToken := []TestTableStruct{
		{
			name: "Test #1. OK. Корректный стандартный запрос",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities",
			"count_verses_on_pages":2,
			"page":1,
			"direction":""
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #2. NOT OK. Неверное direction (page 1, direction previous)",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities",
			"count_verses_on_pages":2,
			"page":1,
			"direction":"previous"
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #3. OK. Запрос на существующую следующую страницу (page 1, direction next)",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities",
			"count_verses_on_pages":2,
			"page":1,
			"direction":"next"
			}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test #4. NOT OK. пустые поля",
			requestBody: `{
			"group":"",
			"song":"",
			"count_verses_on_pages":0,
			"page":0,
			"direction":""
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #5. NOT OK. несуществующая страница",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities",
			"count_verses_on_pages":2,
			"page":999999999,
			"direction":""
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #6. NOT OK. некоректный json (не все поля)",
			requestBody: `{
			"direction":"",
			"page":99999999999
			}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Test #7. NOT OK. некоректный json (есть ненужные поля)",
			requestBody: `{
			"group":"Red Hot Chili Peppers",
			"song":"Dark Necessities",
			"count_verses_on_pages":2,
			"page":4,
			"direction":"next",
			"a":"a"
			}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	return testTableGetToken
}
