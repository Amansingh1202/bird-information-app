<template>
    <div class="layout-content">
        <div class="content-section implementation">
            <div class="card">
                <h5>Basic</h5>
                <Dropdown v-model="selectedCity1" :options="cities" optionLabel="name" optionValue="code" placeholder="Select a City" />

                <h5>Editable</h5>
                <Dropdown v-model="selectedCity2" :options="cities" optionLabel="name" :editable="true"/>

                <h5>Advanced with Templating, Filtering and Clear Icon</h5>
                <Dropdown v-model="selectedCountry" :options="countries" optionLabel="name" :filter="true" placeholder="Select a Country" :showClear="true">
                    <template #value="slotProps">
                        <div class="country-item country-item-value" v-if="slotProps.value">
                            <img src="https://www.primefaces.org/wp-content/uploads/2020/05/placeholder.png" />
                            <div>{{ slotProps.value.name }}</div>
                        </div>
                        <span v-else>
                            {{slotProps.placeholder}}
                        </span>
                    </template>
                    <template #option="slotProps">
                      <div class="country-item">
                        <div class="country-item">
                            <img src="https://www.primefaces.org/wp-content/uploads/2020/05/placeholder.png" />
                            <div>{{ slotProps.option.name }}</div>
                        </div>
                      </div>
                    </template>
                </Dropdown>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            selectedCity1: null,
            selectedCity2: null,
            selectedCountry: null,
            cities: [
                {name: 'New York', code: 'NY'},
                {name: 'Rome', code: 'RM'},
                {name: 'London', code: 'LDN'},
                {name: 'Istanbul', code: 'IST'},
                {name: 'Paris', code: 'PRS'}
            ],
            countries: [
                {name: 'Australia', code: 'AU'},
                {name: 'Brazil', code: 'BR'},
                {name: 'China', code: 'CN'},
                {name: 'Egypt', code: 'EG'},
                {name: 'France', code: 'FR'},
                {name: 'Germany', code: 'DE'},
                {name: 'India', code: 'IN'},
                {name: 'Japan', code: 'JP'},
                {name: 'Spain', code: 'ES'},
                {name: 'United States', code: 'US'}
            ]
        }
    }
}

</script>

<style lang="scss" scoped>
.p-dropdown {
    width: 14rem;
}

.country-item {
    img {
        width: 17px;
        margin-right: 0.5rem;
    }
}
</style>