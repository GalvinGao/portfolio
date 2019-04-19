<template>
  <v-container grid-list-xl>
    <v-layout align-center justify-space-around fill-height v-bind="binding">
      <v-flex xs12 md12 lg5 xl5>
        <div v-if="playerOptions.sources.length !== 0">
          <div v-video-player:myVideoPlayer="playerOptions"></div>
        </div>
        <v-dialog v-else v-model="dialog" hide-overlay persistent width="300">
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
      </v-flex>
      <!--<v-flex xs12 md12 lg5 xl5>-->
      <!--<h5 class="headline">-->
      <!--{{ currentContent.Description }}-->
      <!--</h5>-->
      <!--<h6 class="caption">-->
      <!--{{ formatTime(currentContent.CreatedAt) }}-->
      <!--</h6>-->
      <!--</v-flex>-->
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      currentContent: {},
      dialog: true,
      playerOptions: {
        aspectRatio: '16:9',
        autoplay: false,
        muted: true,
        language: this.$i18n.locale,
        playbackRates: [0.7, 1.0, 1.5, 2.0],
        preload: 'auto',
        responsive: true,
        breakpoints: {
          large: 1440,
          xlarge: 1440,
          huge: 1440
        },

        // sources: [{
        //   type: "video/mp4",
        //   // mp4
        //   src: "http://vjs.zencdn.net/v/oceans.mp4",
        //   // webm
        //   // src: "https://cdn.theguardian.tv/webM/2015/07/20/150716YesMen_synd_768k_vp8.webm"
        // }],
        sources: [],
        poster: '',
        userActions: {
          doubleClick: () => {
            this.pause()
          },
          fullscreenKey: function(event) {
            // override fullscreen to trigger when pressing the v key
            return event.which === 13
          }
        }
      }
    }
  },
  beforeMount() {
    console.log(this.$store.state)
    const that = this

    function cb(data) {
      console.log('data: ', data)

      that.currentContent = data
      that.playerOptions.sources.push({
        type: 'video/mp4',
        src: data.Media
      })
    }

    if (this.$store.state.contentList.length > 0) {
      const content = this.$store.state.contentList.filter(obj => {
        return obj.Permalink === this.$route.params.id
      })[0]
      cb(content)
    } else {
      this.getContent().then(content => {
        cb(content)
      })
    }
  },
  methods: {
    formatTime(str) {
      return this.$moment(str).fromNow()
    },
    async getContent() {
      const response = await this.$axios.get(
        '/projects/' + this.$route.params.id
      )
      return response.data
    }
  },
  computed: {
    binding() {
      const binding = {}

      // if (this.$vuetify.breakpoint.mdAndDown) binding.column = true

      return binding
    }
  }
}
</script>
<style>
.videojs {
  min-height: 640px;
}
.video-js button {
  font-size: 12px;
}
.vjs-playback-rate .vjs-playback-rate-value {
  line-height: 2.5em;
}
.video-js .vjs-big-play-button {
  height: auto;
  width: auto;
  border: none;
  background: none;
  margin-left: -0.5em;
  margin-top: -0.6em;
  top: 50%;
  left: 50%;
}
.vjs-slider-horizontal .vjs-volume-level:before {
  top: -0.333333333333333em;
}
</style>
