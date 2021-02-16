<template>
  <div>
    <Card style="width: 100vw; margin: 2em">
      <template #title>
        Get Random Fish
      </template>
      <template #subtitle>
        Get 3 Random Fish from all the available fish.
      </template>
      <template #content>
        <span>Endpoint</span>http://localhost:3000/api/getRandomFish
        <Divider />
        <span>HTTP Method</span>GET
        <Divider />
        <span>Content-Type</span>application/json
        <Divider />
        <span>Response Format</span>
        <div>
          <pre>
      [{
        "Habitat":""
        "Species Illustration Photo":{
          "Src":""
          "Alt":""
          "Title":""
        }
	      "Location":""
	      "Population":""
	      "Species":""
	      "Color":""
	      "Fat":""
	      "Health":""
	      "Taste":""
      }]
     </pre
          >
        </div>
      </template>
    </Card>
    <Card style="width: 100vw; margin: 2em">
      <template #title>
        Get Fish By Name
      </template>
      <template #subtitle>
        Get fish information by the name of the fish.
      </template>
      <template #content>
        <span>Endpoint</span>https://www.fishwatch.gov/api/species/?name={
        fishName }
        <Divider />
        <span>HTTP Method</span>GET
        <Divider />
        <span>Content-Type</span>application/json
        <Divider />
        <span>Response Format</span>
        <div>
          <pre>
      [{
        "Habitat":""
        "Species Illustration Photo":{
          "Src":""
          "Alt":""
          "Title":""
        }
	      "Location":""
	      "Population":""
	      "Species":""
	      "Color":""
	      "Fat":""
	      "Health":""
	      "Taste":""
      }]
     </pre
          >
        </div>
      </template>
    </Card>
  </div>
</template>

<script>
export default {
  components: {}
};
</script>

<style scoped>
span {
  color: blue;
  font-weight: 600;
  margin-right: 20px;
}
pre {
  background-color: black;
  color: white;
  padding: 10px;
}
div.p-card-body.p-component {
  background-color: gray;
}
</style>
