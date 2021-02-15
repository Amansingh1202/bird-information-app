<template>
  <div>
    <random-fish-gallery />
    <search-fish />
  </div>
</template>

<script>
import RandomFishGallery from "../components/RandomFishGallery.vue";
import SearchFish from "../components/SearchFish.vue";
export default {
  components: {
    RandomFishGallery,
    SearchFish
  }
};
</script>
