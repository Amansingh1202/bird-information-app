<template>
  <div id="container">
    <Navbar id="head1"></Navbar>
    <router-view id="body1"></router-view>
    <footer-component id="foot1" />
  </div>
</template>
<script>
import FooterComponent from "./components/FooterComponent.vue";
import Navbar from "./components/Navbar";
export default {
  components: {
    Navbar,
    FooterComponent
  }
};
</script>
<style lang="scss">
@import "../App.scss";
</style>
