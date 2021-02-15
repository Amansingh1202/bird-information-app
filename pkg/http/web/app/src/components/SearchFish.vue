<template>
  <div class="searchFish">
    <span class="p-input-icon-left">
      <i class="pi pi-search" />
      <InputText type="text" v-model="value" placeholder="Search" /> </span
    ><br />
    <button :disabled="!value" v-on:click="getFish">Submit</button>
    <div v-if="isCalled" class="p-grid">
      <div class="p-col">
        <div class="box">
          <div class="content-section implementation">
            <Card style="width: 30em;;height:75vh">
              <template #header>
                <img
                  :src="fishData[0]['Species Illustration Photo'].src"
                  :alt="fishData[0]['Species Illustration Photo'].alt"
                  style="height: 10rem"
                />
              </template>
              <template #title>
                {{ removeHTMLTags(fishData[0]["Species Name"]) }}
              </template>
              <template #subtitle>
                {{ removeHTMLTags(fishData[0]["Color"]) }}
              </template>
              <template #content>
                <h3>Health Benefits</h3>
                <p>{{ removeHTMLTags(fishData[0]["Health Benefits"]) }}</p>
                <h3>Taste</h3>
                <p>{{ removeHTMLTags(fishData[0]["Taste"]) }}</p>
              </template>
            </Card>
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
      value: null,
      fishData: [],
      isCalled: false
    };
  },
  methods: {
    getFish() {
      fetch("http://localhost:3000/api/getFish?name=" + this.value)
        .then(res => res.json())
        .then(res => {
          this.isCalled = false;
          this.fishData.push(res["message"][0]);
          this.isCalled = true;
        });
    },
    removeHTMLTags(content) {
      content = content.replace(/(<([^>]+)>)/gi, "");
      return content.replace("&nbsp;", "");
    }
  }
};
</script>

<style lang="scss" scoped>
.aligned {
  margin-left: 37vw;
}
.box {
  margin-left: -150px;
}
.searchFish {
  margin-left: 45%;
  margin-bottom: 20px;
}
h5 {
  font-size: 1.6em;
}
button {
  margin-top: 10px;
  color: black;
  background-color: lime;
  border: none;
  font-size: 1.3em;
  padding: 5px 10px;
  border-radius: 4px;
  margin-bottom: 20px;
}
button:hover {
  color: gray;
  background-color: greenyellow;
}
button:disabled {
  background-color: lightgray;
  color: black;
}
</style>
