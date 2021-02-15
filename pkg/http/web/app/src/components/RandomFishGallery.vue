<template>
  <card-component :loaded="loaded" :fishData="fishData" />
</template>

<script>
import CardComponent from "./CardComponent.vue";
export default {
  data() {
    return {
      fishData: [],
      loaded: false
    };
  },
  components: {
    CardComponent
  },
  created() {
    var _this = this;
    function getData() {
      fetch("http://localhost:3000/api/getRandomFish")
        .then(res => res.json())
        .then(response => {
          response = response["message"];
          _this.fishData.push(response[0]);
          _this.fishData.push(response[1]);
          _this.fishData.push(response[2]);
          _this.loaded = true;
        });
    }
    getData();
  }
};
</script>
