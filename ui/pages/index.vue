<template>
  <v-container grid-list-xl fluid>
    <v-layout row wrap>
      <v-flex xs12>
        <h3 class="display-1">{{ $t('category.projects') }}</h3>
      </v-flex>

      <v-flex
        v-for="(content, index) in contentStream"
        :key="index"
        xs12
        sm6
        lg4
      >
        <div @click="enterPage(content)">
          <v-hover>
            <v-card
              dark
              :class="`elevation-${hover ? 8 : 2}`"
              slot-scope="{ hover }"
              nuxt
              ripple
              :to="`project/${content.Permalink}`"
            >
              <v-img :src="content.Thumbnail" aspect-ratio="2.5">
                <v-expand-transition>
                  <div
                    v-if="hover"
                    class="d-flex transition-fast-in-fast-out black darken-2 v-card--reveal display-4 white--text"
                    style="height: 100%;"
                  >
                    <v-icon large>chevron_right</v-icon>
                  </div>
                </v-expand-transition>
              </v-img>
              <v-card-title>
                <div>
                  <div class="headline">{{ content.Name }}</div>
                  <span class="grey--text">{{ content.Description }}</span>
                </div>
              </v-card-title>
            </v-card>
          </v-hover>
        </div>
      </v-flex>

      <v-dialog
        transition="fade-transition"
        v-if="contentStream.length === 0"
        v-model="dialog"
        hide-overlay
        persistent
        width="300"
      >
        <v-card dark>
          <v-card-text>
            {{ $t('meta.loading') }}
            <v-progress-linear
              indeterminate
              color="white"
              class="mb-1"
            ></v-progress-linear>
          </v-card-text>
        </v-card>
      </v-dialog>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  components: {},
  beforeMount() {
    this.getContent().then(content => {
      this.contentStream = content.records
      this.$store.commit('setContentList', content.records)
      console.log(content.records)
    })
  },
  methods: {
    async getContent() {
      const response = await this.$axios.get('/projects', {
        params: {
          page: this.currentPage,
          limit: this.pageLimit
        }
      })
      return response.data
    },
    enterPage(content) {
      let name = content.Name
      console.log('enter page', name)
      this.$store.commit('setExtendedTitle', name)
    }
  },
  data() {
    return {
      dialog: true,
      pageLimit: 20,
      currentPage: 1,
      contentStream: []
    }
  }
}
</script>
<style>
.v-card--reveal {
  align-items: center;
  bottom: 0;
  justify-content: center;
  opacity: 0.8;
  position: absolute;
  width: 100%;
}
.headline {
  margin-bottom: 0.5rem;
}
.v-toolbar__extension {
  height: 48px !important;
}
</style>
