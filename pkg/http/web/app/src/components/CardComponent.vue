<template>
  <div>
    <div class="p-grid">
      <div class="p-col">
        <div class="box">
          <div class="layout-content">
            <div class="content-section implementation">
              <MyLoader v-if="getMyData" />
              <Card v-else style="width: 25em">
                <template #header>
                  <img
                    :src="getMyData[0]['Images'].Src"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  Advanced Card
                </template>
                <template #subtitle>
                  Card subtitle
                </template>
                <template #content>
                  <p>
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Inventore sed consequuntur error repudiandae numquam
                    deserunt quisquam repellat libero asperiores earum nam
                    nobis, culpa ratione quam perferendis esse, cupiditate neque
                    quas!
                  </p>
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
              <MyLoader v-if="getMyData" />
              <Card v-else style="width: 25em">
                <template #header>
                  <img
                    :src="getMyData[1]['Images'].Src"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  Advanced Card
                </template>
                <template #subtitle>
                  Card subtitle
                </template>
                <template #content>
                  <p>
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Inventore sed consequuntur error repudiandae numquam
                    deserunt quisquam repellat libero asperiores earum nam
                    nobis, culpa ratione quam perferendis esse, cupiditate neque
                    quas!
                  </p>
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
              <MyLoader v-if="getMyData" />
              <Card style="width: 25em" v-else>
                <template #header>
                  <img
                    :src="getMyData[2]['Images'].Src"
                    style="height: 10rem"
                  />
                </template>
                <template #title>
                  Advanced Card
                </template>
                <template #subtitle>
                  Card subtitle
                </template>
                <template #content>
                  <p>
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Inventore sed consequuntur error repudiandae numquam
                    deserunt quisquam repellat libero asperiores earum nam
                    nobis, culpa ratione quam perferendis esse, cupiditate neque
                    quas!
                  </p>
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
  }
};
</script>
<style lang="scss" scoped>
p {
  line-height: 1.5;
  margin: 0;
}
</style>
