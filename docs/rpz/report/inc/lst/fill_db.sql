CREATE TABLE json_table (
                            id int PRIMARY KEY,
                            airtableId text,
                            slug text,
                            name text,
                            workplace text,
                            description text,
                            about text,
                            competencies text,
                            experience text,
                            price text,
                            menteeCount int,
                            photo json,
                            photo_url text,
                            tags text[],
                            sortOrder int,
                            is_visible bool,
                            sponsors text,
                            calendarType text
);

WITH json (doc) AS (values ('[
  {
    "id": 336,
    "airtableId": "recBcoVPUSaOKLEqe",
    "slug": "kseniia-pomogaeva-336",
    "name": "Ксения Помогаева",
    "job": "Методолог IT обучения",
    "workplace": "Тинькофф Банк, сертифицированный скрам-мастер",
    "description": "Выстраивание стратегии личного и командного обучения в зависимости от целей (как hard, так и soft), помощь в настройке процессов agile-команд, помощь в реализации программ менторинга, наставничества и обучения внутри компании.",
    "about": "",
    "competencies": "",
    "experience": "2-5",
    "price": "По договоренности",
    "menteeCount": 0,
    "photo": {
      "id": "attBPUGR0pXxjGdYw",
      "width": 709,
      "height": 1057,
      "url": "https://dl.airtable.com/.attachments/d8607779d3ada78c9fad792ed2ece2ea/968154fa/photo5287487268799493337.jpg?ts=1652723915\u0026userId=usrW6Ciyt4Mp0Nk6M\u0026cs=40ef2657e89cfca1",
      "filename": "photo5287487268799493337.jpg",
      "size": 193224,
      "type": "image/jpeg",
      "thumbnails": {
        "small": {
          "url": "https://dl.airtable.com/.attachmentThumbnails/568b737de7b03ef5bc2b87f0076e2911/2a0499a6?ts=1652723915\u0026userId=usrW6Ciyt4Mp0Nk6M\u0026cs=46bbf92159b35348",
          "width": 24,
          "height": 36
        },
        "large": {
          "url": "https://dl.airtable.com/.attachmentThumbnails/318c8a94e2a386ef8b4a2aea9954b232/cdab9fc6?ts=1652723915\u0026userId=usrW6Ciyt4Mp0Nk6M\u0026cs=cd53673135752e05",
          "width": 512,
          "height": 763
        },
        "full": {
          "url": "https://dl.airtable.com/.attachmentThumbnails/07fd1c6d9bb51efe25623d47b84531f3/bcce4de0?ts=1652723915\u0026userId=usrW6Ciyt4Mp0Nk6M\u0026cs=94784d2481c820d8",
          "width": 3000,
          "height": 3000
        }
      }
    },
    "photo_url": "https://dl.airtable.com/.attachments/d8607779d3ada78c9fad792ed2ece2ea/968154fa/photo5287487268799493337.jpg",
    "tags": [
      "Agile",
      "HR",
      "Product Management",
      "Team Lead/Management",
      "Другое"
    ],
    "sortOrder": 4,
    "is_visible": true,
    "sponsors": "none",
    "calendarType": "none"
  }
]'::json))
INSERT INTO json_table
    (id, airtableId, slug, name, workplace,
    description, about, competencies,
     experience, price, menteeCount,
    photo, photo_url, tags, sortOrder,
     is_visible, sponsors, calendarType)
SELECT
    id, airtableId, slug, name, workplace,
    description, about, competencies,
    experience, price, menteeCount,
    photo, photo_url, tags,
    sortOrder, is_visible, sponsors,
    calendarType from json l
cross join lateral
    json_populate_recordset(null::json_table, doc) as p;


INSERT INTO users (first_name, last_name, nickname,
                   avatar, about, is_searchable)
SELECT (string_to_array(p.name, ' '))[1],
       (string_to_array(p.name, ' '))[2],
       p.slug, p.photo_url,
       concat(p.about, ' ', p.description), p.is_visible
From json_table as p;
INSERT INTO skills (name)
SELECT DISTINCT UNNEST(tags) FROM json_table ON CONFLICT DO NOTHING;

INSERT INTO users_skills (user_id, skill_name) (
    SELECT u.id, UNNEST(json_table.tags) from users as u
            JOIN json_table on json_table.slug = u.nickname
)
