<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <div class="text-center">
        <logo />
        <vuetify-logo />
      </div>
      <v-card>
        <v-card-text>
          <p>
            大事なのは自分の心に素直になることだ。
          </p>
          <div class="text-xs-right">
            <em><small>&mdash; スティーブ・ジョブズ</small></em>
          </div>
        </v-card-text>
        <!-- <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire">
            Continue
          </v-btn>
        </v-card-actions> -->
      </v-card>
      <v-card>
        <v-card-text>
          <p>
            大事なのは自分の心に素直になることだ。
          </p>
          <div class="text-xs-right">
            <em><small>&mdash; スティーブ・ジョブズ</small></em>
          </div>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo
  }
}
</script>
