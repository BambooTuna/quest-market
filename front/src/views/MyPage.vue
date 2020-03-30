<template>
  <div class="my-page">
    <Authentication>
      <h2>自分の出品一覧</h2>
      <PrivateProductsTable :params="this.$route.query"></PrivateProductsTable>
    </Authentication>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Authentication from '@/components/Authentication.vue'
import PrivateProductsTable from '@/components/PrivateProductsTable.vue'
@Component({
  components: {
    Authentication, PrivateProductsTable
  }
})

export default class MyPage extends Vue {
}
</script>

<style scoped>

</style>
