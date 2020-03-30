<template>
  <div class="linecooperation">
    リダイレクトされました
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import API from '@/lib/RestAPI'

@Component
export default class LineCooperation extends Vue {
    private api: API = new API()

    created (): void {
      this.api.lineCooperationSignin(this.$route.query)
        .then(() => {
          window.location.href = '/'
          alert('Lineでログインしました')
        })
        .catch(e => {
          window.location.href = '/'
          alert(e.toString())
        })
    }
}
</script>
