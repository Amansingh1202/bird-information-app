<template>
  <code>
    <pre>
            {{ fishData }}
        </pre
    >
  </code>
</template>

<script>
export default {
  data() {
    return {
      fishData: []
    };
  },
  created() {
    var _this = this;
    function getData() {
      fetch("http://localhost:3000/api/getAll")
        .then(response => response.text())
        .then(response => {
          _this.fishData = response;
          console.log(response);
        });
    }
    getData();
  }
};
</script>
