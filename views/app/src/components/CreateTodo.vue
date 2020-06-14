<template>
  <div class="ui basic content center aligned segment">
    <button class="ui basic button icon" v-on:click="openForm" v-show="!isCreating">
      <i class="plus icon"></i>
    </button>
    <div class="ui centered card" v-show="isCreating">
      <div class="content">
        <div class="ui form">
          <div class="field">
            <label>Title</label>
            <input v-model="title" type="text" ref="title" defaultValue @keyup.enter="sendForm()" />
          </div>
          <div class="field">
            <label>Description</label>
            <input
              v-model="description"
              type="text"
              ref="project"
              defaultValue
              @keyup.enter="sendForm()"
            />
          </div>
          <div class="field">
            <label>URL</label>
            <input v-model="url" type="text" ref="project" defaultValue @keyup.enter="sendForm()" />
          </div>
          <div class="ui two button attached buttons">
            <button class="ui basic blue button" v-on:click="sendForm()">Create</button>
            <button class="ui basic red button" v-on:click="closeForm">Cancel</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      description: "",
      url: "",
      isCreating: false
    };
  },
  methods: {
    openForm() {
      this.isCreating = true;
    },
    closeForm() {
      this.isCreating = false;
    },
    sendForm() {
      if (this.title.length > 0) {
        const title = this.title;
        const description = this.description;
        const url = this.url;
        this.$emit("create-todo", {
          title,
          description,
          url,
          done: false
        });
        this.title = "";
        this.description = "";
        this.url = "";
      }
      this.isCreating = false;
    }
  }
};
</script>