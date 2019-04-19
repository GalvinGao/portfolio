<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-tile
          v-for="(item, i) in menuItems"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title :v-text="item.title" />
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar :clipped-left="clipped" fixed app v-bind:extended="!normalTitle">
      <v-toolbar-side-icon @click="drawer = !drawer"></v-toolbar-side-icon>

      <v-toolbar-title v-html="$t('meta.title')"></v-toolbar-title>
      <template v-if="!normalTitle" v-slot:extension>
        <v-toolbar-title
          v-text="extendTitle"
          class="subheading"
        ></v-toolbar-title>
      </template>

      <v-spacer></v-spacer>
      <v-menu bottom left>
        <template v-slot:activator="{ on }">
          <v-btn dark icon v-on="on">
            <v-icon>translate</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-tile
            v-for="(item, i) in translationItems"
            :key="i"
            @click="setTranslation(item.code)"
            ripple
          >
            <v-list-tile-title
              >{{ item.name }}
              <v-icon v-if="locale === item.code"
                >done</v-icon
              ></v-list-tile-title
            >
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar>
    <v-content>
      <v-container>
        <v-breadcrumbs :items="breadcrumbs">
          <template v-slot:divider>
            <v-icon>chevron_right</v-icon>
          </template>
          <template v-slot:item="props">
            <NuxtLink
              ripple
              :to="props.item.href"
              class="v-breadcrumbs__item grey--text text--lighten-4"
              >{{ props.item.text }}</NuxtLink
            >
          </template>
        </v-breadcrumbs>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer app style="padding-left: 1em;" transition="slide-x-transition">
      <span
        >&copy; {{ new Date().getFullYear() }} {{ $t('meta.copyright') }}</span
      >
    </v-footer>
  </v-app>
</template>

<script>
export default {
  beforeMount() {
    this.updateTitle()
    setTimeout(() => {
      this.updateBreadcrumb()
    }, 0)
    this.locale = this.$i18n.locale
    console.log('breadcrumbs: ', this.breadcrumbs)
  },
  data() {
    return {
      clipped: false,
      drawer: false,
      menuItems: [
        {
          icon: 'view_module',
          title: 'Projects',
          to: '/'
        },
        {
          icon: 'info',
          title: 'About',
          to: '/about'
        }
      ],
      translationItems: [
        {
          name: 'English',
          code: 'en'
        },
        {
          name: '中文',
          code: 'zh'
        }
      ],
      miniVariant: true,
      locale: null,
      normalTitle: true,
      extendTitle: this.$store.state.extendedTitle,
      breadcrumbs: []
    }
  },
  methods: {
    setTranslation(lang) {
      if (this.$i18n.locale === lang) return false
      document.querySelector('.application--wrap').style.opacity = 0
      setTimeout(() => {
        this.$i18n.locale = lang
        this.$store.commit('setLang', lang)
        this.locale = lang
      }, 175)

      setTimeout(() => {
        document.querySelector('.application--wrap').style.opacity = 1
      }, 175)
    },
    updateTitle() {
      setTimeout(() => {
        this.normalTitle = this.$route.path.indexOf('project') < 0

        if (this.$store.state.extendedTitle) {
          this.extendTitle = this.$store.state.extendedTitle
        } else {
          this.getProjectByPermalink(this.$route.params.id).then(data => {
            this.extendTitle = data.Name
          })
        }
      }, 0)
    },
    async getProjectByPermalink(query) {
      const response = await this.$axios.get('/projects/' + query)
      return response.data
    },
    pushBreadcrumb(content) {
      this.breadcrumbs.push({
        text: content.Name,
        disabled: false,
        href: content.Permalink
      })
    },
    updateBreadcrumb() {
      let path = this.$route.path.split('/').slice(1)
      console.log('path: ', path, this.$route)

      // default
      this.breadcrumbs = [
        {
          text: this.$t('category.home'),
          disabled: false,
          href: '/'
        }
      ]

      if (path[0] === 'project') {
        let contentListFromStore = this.$store.state.contentList
        let projectId = path[1]
        if (contentListFromStore.length === 0) {
          this.getProjectByPermalink(projectId).then(data => {
            this.pushBreadcrumb(data)
          })
        } else {
          this.pushBreadcrumb(
            contentListFromStore.filter(obj => {
              return obj.Permalink === projectId
            })[0]
          )
        }
      } else if (path[0] === 'about') {
        this.breadcrumbs.push({
          text: this.$t('category.about'),
          disabled: false,
          href: path[0]
        })
      }
    }
  },
  watch: {
    $route: ['updateTitle', 'updateBreadcrumb']
  }
}
</script>

<style>
.application--wrap {
  transition: all 225ms cubic-bezier(0, 0, 0.2, 1);
}
</style>
