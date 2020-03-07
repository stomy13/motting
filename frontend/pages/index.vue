<template>
  <div class="text-center">
    <v-row>
      <logo />
      <vuetify-logo />
      <push-time />
    </v-row>

    <phrase />
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'
import Phrase from '~/components/Phrase.vue'
import PushTime from '~/components/PushTime.vue'

export default {
  components: {
    Logo,
    VuetifyLogo,
    Phrase,
    PushTime
  }
}
</script>
