<template>
  <div>
    <random-fish-gallery />
    <FooterComponent />
  </div>
</template>

<script>
import FooterComponent from "../components/FooterComponent";
import RandomFishGallery from "../components/RandomFishGallery.vue";
export default {
  components: {
    FooterComponent,
    RandomFishGallery
  }
};
</script>
