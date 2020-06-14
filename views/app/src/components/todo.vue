<template>
  <div class="ui centered card">
    <div class="content" v-show="!isEditing">
      <div v-if="todo.Url">
        <a :href="todo.Url" target="_blank">
          <div class="header">{{ todo.Title }}</div>
        </a>
      </div>
      <div v-else>
        <div class="header">{{ todo.Title }}</div>
      </div>

      <div class="meta">{{ todo.Description }}</div>
      <div class="extra content">
        <span class="right floated edit icon" v-on:click="showForm">
          <i class="edit icon"></i>
        </span>
        <span class="right floated trash icon" v-on:click="deleteTodo(todo)">
          <i class="trash icon"></i>
        </span>
      </div>
    </div>
    <div class="content" v-show="isEditing">
      <div class="ui form">
        <form>
          <div class="field">
            <input type="hidden" name="Id" :value="todo.Id" />
            <label>Title</label>
            <input type="text" name="Title" :value="todo.Title" />
          </div>
          <div class="field">
            <label>Description</label>
            <input type="text" name="Description" :value="todo.Description" />
          </div>
          <div class="field">
            <label>URL</label>
            <input type="text" name="Url" :value="todo.Url" />
          </div>
          <div class="ui two button attached buttons">
            <button class="ui basic blue button" v-on:click="saveForm">
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
    <div
      v-on:click="completeTodo(todo)"
      class="ui bottom attached green basic button"
      v-show="!isEditing && todo.Done"
      disabled
    >
      Completed
    </div>
    <div
      v-on:click="completeTodo(todo)"
      class="ui bottom attached red basic button"
      v-show="!isEditing && !todo.Done"
    >
      Pending
    </div>
  </div>
</template>

<script type="text/javascript">
export default {
  props: ["todo"],
  data() {
    return {
      isEditing: false,
      title: "",
      description: "",
      url: "",
      id: null,
    };
  },
  methods: {
    showForm() {
      this.isEditing = true;
    },
    saveForm(e) {
      this.isEditing = false;
      e.preventDefault();
      const inputs = e.target.form.getElementsByTagName("input");
      var td = {};
      inputs.forEach((input) => (td[input.name] = input.value));
      this.$emit("update-todo", td);
    },
    completeTodo(todo) {
      todo.done = !todo.done;
    },
    deleteTodo(todo) {
      this.$emit("delete-todo", todo);
    },
  },
};
</script>
