<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>URL Shortener Microservice</title>
</head>
<body>

<h1>Внимание!!!</h1>
<h2>Это тестовый код. Он имеет множество ошибок и опасных моментов.</h2>

<h2>Описание микросервиса для сокращения URL</h2>
<p>Этот микросервис создан в соответствии с <strong>техническим заданием (ТЗ)</strong>, чтобы обеспечить функционал сокращения URL-адресов. Основная цель микросервиса - создание коротких URL, которые могут быть использованы вместо длинных URL-адресов.</p>

<h3>Основные требования</h3>
<ul>
    <li>Длина сокращенного URL-адреса должна быть <em>как можно короче</em>.</li>
    <li>Сокращенный URL может содержать цифры (<code>0-9</code>) и буквы (<code>a-z</code>, <code>A-Z</code>).</li>
</ul>

<h3>Эндпоинты</h3>

<h4>1. Создание сокращенного URL</h4>
<ul>
    <li><strong>Метод</strong>: POST</li>
    <li><strong>URL</strong>: <a href="http://localhost:8080/">http://localhost:8080/</a></li>
    <li><strong>Request (body)</strong>: <br>
    <code>http://cjdr17afeihmk.biz/123/kdni9/z9d112423421</code></li>
    <li><strong>Response</strong>: <br>
    <code>http://localhost:8080/qtj5opu</code></li>
</ul>

<h4>2. Получение полного URL по сокращенному URL</h4>
<ul>
    <li><strong>Метод</strong>: GET</li>
    <li><strong>URL</strong>: <a href="http://localhost:8080/qtj5opu">http://localhost:8080/qtj5opu</a></li>
    <li><strong>Request (url query)</strong>: <br>
    <code>http://localhost:8080/qtj5opu</code></li>
    <li><strong>Response (body)</strong>: <br>
    <code>http://cjdr17afeihmk.biz/123/kdni9/z9d112423421</code></li>
</ul>

<h3>Хранение данных</h3>
<p>Микросервис имеет возможность хранить информацию в <em>памяти</em> или в базе данных <em>PostgreSQL</em> в зависимости от флага запуска.</p>

<h3>Язык программирования</h3>
<p>Этот микросервис написан на языке программирования <em>Go</em> (Golang), обеспечивая <em>быстродействие</em> и <em>эффективное использование ресурсов</em>.</p>

<h3>Зависимости</h3>
<p>Для работы микросервиса требуется наличие <em>Go</em> версии X.X.X и установленных зависимостей, указанных в файле <code>go.mod</code>.</p>

<h3>Использование</h3>
<p>Чтобы запустить микросервис, выполните следующие шаги:</p>
<ol>
    <li>Клонируйте репозиторий на локальную машину.</li>
    <li>Убедитесь, что все зависимости установлены с помощью <code>go mod tidy</code>.</li>
    <li>Используйте <code>go build</code> для сборки проекта.</li>
    <li>Запустите исполняемый файл, указав необходимые параметры (например, флаг хранения данных).</li>
</ol>

<h3>Автор</h3>
<p>Если у вас есть какие-либо вопросы или предложения по улучшению функционала, не стесняйтесь создавать <em>issues</em> или отправлять <em>pull requests</em> в репозиторий проекта. <em>Ваш вклад приветствуется!</em></p>

<p><strong>Автор</strong>: Тихомиров Максим Русланович<br>
<strong>Контактная информация</strong>: <a href="mailto:awesome.gail@yandex.ru">awesome.gail@yandex.ru</a>, <a href="https://t.me/Tichomirov2003">Telegram</a></p>

<!-- Language Switcher -->
<a href="#english-version">English Version</a> | <a href="#русская-версия">Русская версия</a>

<hr>

<!-- English Version -->
<h2 id="english-version">Attention!!!</h2>
<h2>This is test code. It has many errors and dangerous issues.</h2>

<h2>Description of the URL Shortener Microservice</h2>
<p>This microservice is created according to the <strong>technical specification (TS)</strong> to provide URL shortening functionality. The main goal of the microservice is to create short URLs that can be used instead of long URLs.</p>

<h3>Main Requirements</h3>
<ul>
    <li>The length of the shortened URL should be <em>as short as possible</em>.</li>
    <li>The shortened URL may contain digits (<code>0-9</code>) and letters (<code>a-z</code>, <code>A-Z</code>).</li>
</ul>

<h3>Endpoints</h3>

<h4>1. Create a shortened URL</h4>
<ul>
    <li><strong>Method</strong>: POST</li>
    <li><strong>URL</strong>: <a href="http://localhost:8080/">http://localhost:8080/</a></li>
    <li><strong>Request (body)</strong>: <br>
    <code>http://cjdr17afeihmk.biz/123/kdni9/z9d112423421</code></li>
    <li><strong>Response</strong>: <br>
    <code>http://localhost:8080/qtj5opu</code></li>
</ul>

<h4>2. Get the full URL by the shortened URL</h4>
<ul>
    <li><strong>Method</strong>: GET</li>
    <li><strong>URL</strong>: <a href="http://localhost:8080/qtj5opu">http://localhost:8080/qtj5opu</a></li>
    <li><strong>Request (url query)</strong>: <br>
    <code>http://localhost:8080/qtj5opu</code></li>
    <li><strong>Response (body)</strong>: <br>
    <code>http://cjdr17afeihmk.biz/123/kdni9/z9d112423421</code></li>
</ul>

<h3>Data Storage</h3>
<p>The microservice can store information either <em>in memory</em> or in a <em>PostgreSQL database</em> depending on the startup flag.</p>

<h3>Programming Language</h3>
<p>This microservice is written in the <em>Go</em> programming language (Golang), ensuring <em>high performance</em> and <em>efficient resource usage</em>.</p>

<h3>Dependencies</h3>
<p>The microservice requires <em>Go</em> version X.X.X and the dependencies listed in the <code>go.mod</code> file to run.</p>

<h3>Usage</h3>
<p>To run the microservice, follow these steps:</p>
<ol>
    <li>Clone the repository to your local machine.</li>
    <li>Ensure all dependencies are installed using <code>go mod tidy</code>.</li>
    <li>Use <code>go build</code> to build the project.</li>
    <li>Run the executable, specifying the necessary parameters (e.g., data storage flag).</li>
</ol>

<h3>Author</h3>
<p>If you have any questions or suggestions for improving functionality, feel free to create <em>issues</em> or submit <em>pull requests</em> to the project repository. <em>Your contribution is welcome!</em></p>

<p><strong>Author</strong>: Tikhomirov Maxim Ruslanovich<br>
<strong>Contact Information</strong>: <a href="mailto:awesome.gail@yandex.ru">awesome.gail@yandex.ru</a>, <a href="https://t.me/Tichomirov2003">Telegram</a></p>

</body>
</html>
