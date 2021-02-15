<template>
  <div>
    <div class="p-grid">
      <div class="p-col">
        <div class="box">
          <div class="layout-content">
            <div class="content-section implementation">
              <MyLoader v-if="!getMyData" />
              <Card v-else style="width: 25em">
                <template #header>
                  <img
                    :src="getMyData[0]['Species Illustration Photo'].src"
                    :alt="getMyData[1]['Species Illustration Photo'].alt"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  {{ removeHTMLTags(getMyData[0]["Species Name"]) }}
                </template>
                <template #subtitle>
                  {{ removeHTMLTags(getMyData[0]["Color"]) }}
                </template>
                <template #content>
                  <h3>Health Benefits</h3>
                  <p>{{ removeHTMLTags(getMyData[0]["Health Benefits"]) }}</p>
                  <h3>Taste</h3>
                  <p>{{ removeHTMLTags(getMyData[0]["Taste"]) }}</p>
                </template>
              </Card>
            </div>
          </div>
        </div>
      </div>
      <div class="p-col">
        <div class="box">
          <div class="layout-content">
            <div class="content-section implementation">
              <MyLoader v-if="!getMyData" />
              <Card v-else style="width: 25em">
                <template #header>
                  <img
                    :src="getMyData[1]['Species Illustration Photo'].src"
                    :alt="getMyData[1]['Species Illustration Photo'].alt"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  {{ removeHTMLTags(getMyData[1]["Species Name"]) }}
                </template>
                <template #subtitle>
                  {{ removeHTMLTags(getMyData[1]["Color"]) }}
                </template>
                <template #content>
                  <h3>Health Benefits</h3>
                  <p>{{ removeHTMLTags(getMyData[1]["Health Benefits"]) }}</p>
                  <h3>Taste</h3>
                  <p>{{ removeHTMLTags(getMyData[1]["Taste"]) }}</p>
                </template>
              </Card>
            </div>
          </div>
        </div>
      </div>
      <div class="p-col">
        <div class="box">
          <div class="layout-content">
            <div class="content-section implementation">
              <MyLoader v-if="!getMyData" />
              <Card style="width: 25em" v-else>
                <template #header>
                  <img
                    :src="getMyData[2]['Species Illustration Photo'].src"
                    :alt="getMyData[2]['Species Illustration Photo'].alt"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  {{ removeHTMLTags(getMyData[2]["Species Name"]) }}
                </template>
                <template #subtitle>
                  {{ removeHTMLTags(getMyData[2]["Color"]) }}
                </template>
                <template #content>
                  <h3>Health Benefits</h3>
                  <p>{{ removeHTMLTags(getMyData[2]["Health Benefits"]) }}</p>
                  <h3>Taste</h3>
                  <p>{{ removeHTMLTags(getMyData[2]["Taste"]) }}</p>
                </template>
              </Card>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { MyLoader } from "../components/MyLoader";
export default {
  data() {
    return {};
  },
  props: ["fishData", "loaded"],
  computed: {
    getMyData() {
      if (this.loaded) {
        return this.fishData;
      } else {
        return false;
      }
    }
  },
  components: {
    MyLoader
  },
  methods: {
    removeHTMLTags(content) {
      content = content.replace(/(<([^>]+)>)/gi, "");
      return content.replace("&nbsp;", "");
    }
  }
};
</script>
<style lang="scss" scoped>
p {
  line-height: 1.5;
  margin: 0;
}
</style>
