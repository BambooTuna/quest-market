<template>
  <div class="home">
    <PublicProductsTable :params="this.$route.query"></PublicProductsTable>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import PublicProductsTable from '@/components/PublicProductsTable.vue'
@Component({
  components: {
    PublicProductsTable
  }
})

export default class Home extends Vue {
}
</script>
