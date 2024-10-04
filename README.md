Тестовое задание в компанию Effective Mobile.<br/>

Описание:<br/>
1)Структура проекта:<br/>
В проекте находятся следующие папки: cmd, config, elements, handlers, pkg.<br/>
cmd - содержит исполняемый файл main.go.<br/>
config - содержит config.yml с данными для работы программы.<br/>
elements - содержит элемент сервер.<br/>
handlers - содержит файл с путями по которым будут идти запросы.<br/>
в pkg находятся следующие папки repository,service,tests.<br/>
В repisitory хранятся скрипты отвечающие за взаимодействие с БД.<br/>
В service вызываются промежуточные функции и функции взаимодействия с БД.<br/>
в tests лежат тесты.<br/>
Логика приложения следующая: запрос -> handler -> service -> repository.<br/>

2)Структура БД:<br/>
В БД 3 таблицы verses(куплеты) songs и groups.<br/>
В songs хранится id, group_id, name, link,release_date.<br/>
В groups id,name.<br/>
В verses id,song_id,verse_number,text.<br/>
Связаны таблицы: verses song_id -> songs id. groups id -> songs groupd_id.<br/>

Краткое описание работы программы:<br/>
В Программе следущие handlers: /create_new_song, /get_all_songs, /info, /delete_song, /change_song, /get_song_by_verses<br/>
В базе данных находятся 5 базовых песен.<br/>
Swagger работает по url: /docs/*any .<br/>
После добавления песни идёт сразу get запрос по /info и возвращается результат /info.<br/>
Пагинация есть на /get_all_songs и /get_song_by_verses. Обязательные параметры page, и count_on_page от них зависит вывод,
они указываются в json. direction принимает next и previous(может быть пустым), работает от введенной page. field в get_all_songs 
Принимает name и group, filter desc и asc. field и filter могут быть пустыми, но если указан параметр в любом из них, 2 тоже должен
быть заполнен<br/>

Формат запроса на /get_all_songs<br/>
{<br/>
  "direction":"",<br/>
  "page":1,<br/>
  "filter":"",<br/>
  "field":"",<br/>
  "count_songs_on_page":2<br/>
}<br/>

Формат запроса на /get_song_by_verses
{
   "group":"Red Hot Chili Peppers",
   "song":"Dark Necessities",
   "count_verses_on_pages":2,
   "page":1,
   "direction":""
}

Docker:<br/>
Чтобы использовать docker нужно прописать docker-compose up находясь в папке effective_mobile. !!!Так как docker не видит postgre если host:localhost  в .env он уже
Настроен на postgre. его можно изменить закомментив строку с postgre и разкоменнтить с localhost<br/>

Tests:<br/>
Тесты лежат в паке tests. в ней два файла в all_test.go запускаются тесты, в test_table.go лежать сами тесты.<br/>
