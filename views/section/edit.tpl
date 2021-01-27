<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">编辑角色</div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form action="/section/edit/{{.Section.Id}}" method="post">
          <div class="form-group">
            <label for="name">版块名称</label>
            <input type="text" id="name" name="name" value="{{.Section.Name}}" class="form-control">
          </div>
          <button type="submit" class="btn btn-sm btn-default">保存</button>
        </form>
      </div>
    </div>
  </div>
</div>
