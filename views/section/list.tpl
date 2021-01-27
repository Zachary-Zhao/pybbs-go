<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        版块管理
        <a href="/section/add" class="pull-right">添加版块</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .Sections}}
          <tr>
            <td>{{.Id}}</td>
            <td>{{.Name}}</td>
            <td>
              <a href="/section/edit/{{.Id}}" class="btn btn-xs btn-warning">编辑版块</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/section/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>