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
        <div class="field">
          <input type="hidden" name="Id" v-model="id" />
          <input type="hidden" name="Active" v-model="active" />
          <label>Title</label>
          <input type="text" name="Title" v-model="title" />
        </div>
        <div class="field">
          <label>Description</label>
          <input type="text" name="Description" v-model="description" />
        </div>
        <div class="field">
          <label>URL</label>
          <input type="text" name="Url" v-model="url" />
        </div>
        <div class="ui two button attached buttons">
          <button class="ui basic blue button" v-on:click="saveForm">
            Save
          </button>
          <button class="ui basic red button" v-on:click="closeForm">
            Cancel
          </button>
        </div>
      </div>
    </div>
    <div
      v-on:click="completeTodo(todo)"
      class="ui bottom attached green basic button"
      v-show="!isEditing && todo.Active"
      disabled
    >
      Active
    </div>
    <div
      v-on:click="completeTodo(todo)"
      class="ui bottom attached red basic button"
      v-show="!isEditing && !todo.Active"
    >
      Inactive
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
      active: false,
    };
  },
  methods: {
    showForm() {
      this.isEditing = true;
    },

    closeForm() {
      this.isEditing = false;
      this.title = this.todo.Title;
      this.description = this.todo.Description;
      this.url = this.todo.Url;
    },

    saveForm() {
      const Title = this.title;
      const Description = this.description;
      const Url = this.url;
      const Id = this.id;
      // const Active = this.active;
      this.$store.dispatch("tds/updateTodo", {
        Title,
        Description,
        Url,
        Id,
        Active: false,
      });
      this.isEditing = false;
    },

    completeTodo(todo) {
      const Active = !todo.Active;
      const Title = todo.Title;
      const Description = todo.Description;
      const Url = todo.Url;
      const Id = todo.Id;
      this.$store.dispatch("tds/updateTodo", {
        Title,
        Description,
        Url,
        Id,
        Active,
      });
    },

    deleteTodo(todo) {
      this.$emit("delete-todo", todo);
    },
  },

  created() {
    this.title = this.todo.Title;
    this.description = this.todo.Description;
    this.url = this.todo.Url;
    this.id = this.todo.Id;
    this.active = this.todo.Active;
  },
};
</script>
