<template>
  <card-component />
</template>

<script>
import CardComponent from "./CardComponent.vue";
export default {
  data() {
    return {
      fishData: []
    };
  },
  components: {
    CardComponent
  },
  created() {
    // var _this = this;
    function getData() {
      fetch("http://localhost:3000/api/getRandomFish")
        .then(res => res.json())
        .then(response => {
          // console.log(response);
          response = response["message"];
          var message = JSON.parse(response);
          console.log(message);
          // _this.fishData = message[0];
        });
    }
    getData();
  }
};
</script>
