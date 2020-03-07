<template>
  <div class="text-center">
    <logo />
    <vuetify-logo />
    <phrase />
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'
import Phrase from '~/components/Phrase.vue'

export default {
  components: {
    Logo,
    VuetifyLogo,
    Phrase
  }
}
</script>
