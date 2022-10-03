# mini-grep

$ mini-grep [опции] шаблон путь_к_файла


* Опции - эдополнительные параметры, с помощью которых указываются различные настройки поиска и вывода, например количество строк или режим инверсии.
* Шаблон - любая строка ~~или регулярное выражение~~, по которому будет вестись поиск
* Имя файл - файл по которому будет вестись поиск

## Опции

 -A (int)   
 &emsp;&emsp;&emsp;печатать +N строк после совпадения (default 0)  
 -B (int)   
 &emsp;&emsp;&emsp;печатать +N строк до совпадения  (default 0)  
 -C (int)   
&emsp;&emsp;&emsp; печатать ±N строк вокруг совпадения  (default 0)
  
 -c    напечатать количество строк  
 -f    точное совпадение со строкой  
 -i    игнорировать регистр  
 -n    печатать номер строки  
 -v    вместо совпадения, исключать  
 
 ## Пример
  go run .\grep.go "call" data/1.txt  
  go run .\grep.go -A 1 -n -i "CAll" data/1.txt
