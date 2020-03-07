<template>
  <v-row no-gutters>
    <v-col v-for="(phrase, i) in phrases" :key="i" cols="12" sm="4">
      <!-- <v-card v-for="(phrase, i) in phrases" :key="i"> -->
      <v-card>
        <v-card-text>
          <p>
            {{ phrase.text }}
          </p>
          <div class="text-xs-right">
            <em>
              <small>&mdash; {{ phrase.author }}</small>
            </em>
          </div>
        </v-card-text>
        <!-- <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire">
          Continue
          </v-btn>
      </v-card-actions> -->
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  computed: {
    phrases() {
      return this.$store.state.phrases.list
    }
  },
  methods: {
    addPhrases(e) {
      this.$store.commit('phrases/add', e.target.value)
      e.target.value = ''
    }
  }
}
</script>
