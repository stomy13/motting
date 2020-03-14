<template>
  <v-row no-gutters>
    <v-col v-for="(phrase, i) in phrases" :key="i" cols="12" sm="4">
      <v-card>
        <v-card-text>
          <p>
            {{ phrase.Text }}
          </p>
          <div class="text-xs-right">
            <em>
              <small>&mdash; {{ phrase.Author }}</small>
            </em>
          </div>
        </v-card-text>
        <v-btn
          x-small
          absolute
          right
          bottom
          dark
          @click="removePhrase(phrase.ID)"
        >
          <v-icon small dark>mdi-delete</v-icon>
        </v-btn>
      </v-card>
    </v-col>
    <v-btn
      class="mx-2"
      fixed
      fab
      right
      dark
      color="indigo"
      @click="dialog = !dialog"
    >
      <v-icon dark>mdi-plus</v-icon>
    </v-btn>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-text>
          <v-text-field v-model="text" label="Text"></v-text-field>
          <v-text-field v-model="author" label="Author"></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="addPhrase">Add</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      text: '',
      author: ''
    }
  },
  computed: {
    phrases() {
      return this.$store.state.phrases.list
    }
  },
  mounted() {
    this.$store.dispatch('phrases/fetch')
  },
  methods: {
    async addPhrase() {
      await this.$store.dispatch({
        type: 'phrases/post',
        text: this.text,
        author: this.author
      })
      this.dialog = false
      this.text = ''
      this.author = ''
    },
    removePhrase(id) {
      this.$store.dispatch({
        type: 'phrases/delete',
        id
      })
    }
  }
}
</script>
