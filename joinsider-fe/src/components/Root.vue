<template>
    <v-app>
        <v-app-bar app dark>
            <v-app-bar-nav-icon></v-app-bar-nav-icon>
            <v-toolbar-title><router-link to="/">Football Manager</router-link></v-toolbar-title>
        </v-app-bar>
        <v-main>
            <v-container>
                <router-view></router-view>
            </v-container>
        </v-main>
        <v-spacer/>
        <v-footer app>
        </v-footer>
    </v-app>
</template>

<style>
    @import '~vuetify/dist/vuetify.min.css';
</style>
