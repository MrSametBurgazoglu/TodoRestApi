# TodoRestApi
TodoRestAPI with Go,Echo,Gorm,Mysql

### Get All Todos
```
http GET IP/:PORT/todo/
```

### Create New Todo

```
http POST IP/:PORT/todo/ context="context" done:=false
```

### Update Todo By ID

```
http POST IP/:PORT/todo/id context="new context" done:=true
```


### Delete Todo By ID

```
http DELETE IP/:PORT/todo/id
```
