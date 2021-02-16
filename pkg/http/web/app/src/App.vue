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
.p-menubar {
  background: var(--navbar-background);
  border: none;
}
.p-menubar
  .p-menubar-root-list
  > .p-menuitem
  > .p-menuitem-link
  .p-menuitem-text {
  color: var(--navbar-text) !important;
}
</style>
