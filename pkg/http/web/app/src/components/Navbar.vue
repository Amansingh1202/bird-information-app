<template>
  <div class="content-section implementation">
    <div class="card">
      <Menubar :model="items">
        <template #start>
          <img
            alt="logo"
            src="../assets/images/logo.png"
            height="40"
            class="p-mr-2"
          />
        </template>
        <template #end>
          <i
            @click="changeDarkMode()"
            v-if="darkModeOn"
            class="pi pi-moon"
            style="fontSize: 2rem"
          ></i>
          <i
            @click="changeDarkMode()"
            v-else
            class="pi pi-sun"
            style="fontSize: 2rem"
          ></i>
        </template>
      </Menubar>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      darkModeOn: false,
      items: [
        {
          label: "Home",
          to: "/"
        },
        {
          label: "Explore",
          to: "/explore"
        },
        {
          label: "About Me",
          to: "/aboutme"
        },
        {
          label: "API Docs",
          to: "/apidocs"
        }
      ]
    };
  },
  methods: {
    changeDarkMode() {
      this.darkModeOn = !this.darkModeOn;
    }
  }
};
</script>

<style scoped>
.p-menubar {
  background: white;
}
</style>
