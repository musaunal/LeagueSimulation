<template>
    <div>
        <v-row>
            <v-col>
                <v-btn color="primary" @click="distributeFixture">Fixtürü Hazırla</v-btn>
                <v-btn color="primary" @click="playAll">Tüm ligi oynat</v-btn>
                <v-btn color="primary" @click="playOneWeek">Sonraki Haftayı Oynat</v-btn>
                <v-btn color="danger">Ligi Sıfırla</v-btn>
            </v-col>
        </v-row>
        <v-spacer/>
        <v-data-table :headers="table.headers" :loading="table.loading" :items="table.data">
            <template v-slot:top>
                <v-toolbar-title>Takım Listesi</v-toolbar-title>
                <v-divider class="mx-4" inset vertical/>
                <v-spacer/>
                <v-dialog max-width="500px" v-model="dialog.show">
                    <template v-slot:activator="{on, attrs}">
                        <v-btn color="primary" dark v-bind="attrs" v-on="on">
                            Yeni Takım Ekle
                        </v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <span class="headline">Takım Ekle/Düzenle</span>
                        </v-card-title>
                        <v-card-text>
                            <v-container>
                                <v-row>
                                    <v-col cols="12">
                                        <v-text-field v-model="dialog.data.name" label="Takım İsmi"/>
                                    </v-col>
                                </v-row>
                            </v-container>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer/>
                            <v-btn text @click="close">Kapat</v-btn>
                            <v-btn text @click="save">Kaydet</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </template>

            <template v-slot:[`item.name`]="{item}">
                <router-link :to="`/leagues/0/teams/${item.id}`">{{ item.name }}</router-link>
            </template>
            <template v-slot:[`item.actions`]="{item}">
                <v-icon
                        small
                        class="mr-2"
                        @click="editItem(item)">
                    mdi-pencil
                </v-icon>
                <v-icon
                        small
                        @click="deleteItem(item.id)">
                    mdi-delete
                </v-icon>
            </template>
        </v-data-table>
        <v-spacer/>
        <div>
            <v-row>
                <v-col :cols="6">
                    <v-toolbar-title>Fikstür</v-toolbar-title>
                    <v-spacer/>
                    <div dense v-for="(item, week) in _.groupBy(fixture, 'week')" :key="week">
                        <span>{{ week }}.Hafta</span>
                        <v-row dense>
                            <v-col :cols="6" v-for="match in item" :key="match.id">
                                <v-card>
                                    <v-card-text>
                                        <b>{{ match.home_team_name }} vs {{ match.away_team_name }}</b>
                                        <div v-if="match.is_played == 0">
                                            Maç henüz başlamadı...
                                        </div>
                                        <div v-else>
                                            <span>{{ match.home_team_score }} - {{ match.away_team_score }}</span>
                                        </div>
                                        <div>
                                            <span>Hakem: Cüneyt Çakır</span>
                                        </div>
                                    </v-card-text>
                                </v-card>
                            </v-col>
                        </v-row>
                    </div>
                </v-col>
                <v-col :cols="6">
                    <v-toolbar-title>Tahminler</v-toolbar-title>
                    <v-spacer/>
                    <v-row dense>
                        <v-col :cols="12" v-for="data in predict" :key="data.team_id">
                            <v-card>
                                <v-card-text>
                                    <b>{{data.team_name}} -> {{data.rate}}%</b>
                                </v-card-text>
                            </v-card>
                        </v-col>
                    </v-row>
                </v-col>
            </v-row>
        </div>

        <div>
            <v-data-table :headers="scoreBoard.headers" :loading="table.loading" :items="scoreBoard.data"
                          :sort-by="['points', 'goal_diff']"
                          :sort-desc="[true, true]"
            >
                <template v-slot:top>
                    <v-toolbar-title>Puan Durumu</v-toolbar-title>
                </template>
            </v-data-table>
        </div>
        <v-overlay :value="table.loading">
            <v-progress-circular
                    indeterminate
                    size="64"
            ></v-progress-circular>
        </v-overlay>
    </div>
</template>

<script>
    import * as _ from 'lodash';
    export default {
        data() {
            return {
                dialog: {
                    show: false,
                    data: {
                        name: '',
                    }
                },
                table: {
                    loading: false,
                    headers: [
                        {text: 'Takım İsmi', value: 'name'},
                        {text: 'Oynadığı Lig', value: 'league.name'},
                        {text: 'Güncelleme', value: 'updated_at'},
                        {text: 'Aksiyonlar', value: 'actions', sortable: false},
                    ],
                    data: []
                },
                fixture: [],
                scoreBoard: {
                    data: [],
                    headers: [
                        {text: 'Takım İsmi', value: 'team.name'},
                        {text: 'G', value: 'won'},
                        {text: 'B', value: 'drawn'},
                        {text: 'M', value: 'lost'},
                        {text: 'A', value: 'for'},
                        {text: 'Y', value: 'against'},
                        {text: 'AV', value: 'goal_diff'},
                        {text: 'P', value: 'points'},
                    ]
                },
                predict: [],
            }
        },
        computed: {
            leagueId() {
                return this.$route.params.league_id
            },
            _() {
                return _;
            }
        },
        mounted() {
            this.init()
        },
        methods: {
            /**
             * @returns {void}
             */
            async init() {
                await this.refresh();
                this.table.loading = true;
                const {data} = await this.$api.league.get(this.leagueId);
                this.fixture = data.fixture;
                this.scoreBoard.data = data.scoreboard;
                const predict = await this.$api.match.predict(this.leagueId);
                this.predict = predict.data;
                this.table.loading = false;
            },

            /**
             * @returns {void}
             */
            async save() {
                this.dialog.show = false;
                this.table.loading = true;
                if (this.dialog.data.id) {
                    await this.$api.team.update(this.leagueId, this.dialog.data.id, {name: this.dialog.data.name});
                } else {
                    await this.$api.team.create(this.leagueId, {name: this.dialog.data.name})
                }
                await this.close();
                await this.refresh();
            },

            /**
             * @param {object} item
             * @returns {void}
             */
            async editItem(item) {
                this.dialog.data = item;
                this.dialog.show = true;
            },

            /**
             * @param {number} id
             * @returns {void}
             */
            async deleteItem(id) {
                this.table.loading = true;
                await this.$api.team.delete(this.leagueId, id);
                await this.refresh();
            },

            /**
             * @returns {void}
             */
            async close() {
                this.dialog.show = false;
                this.dialog.data = {}
            },

            /**
             * @returns {void}
             */
            async refresh() {
                this.table.loading = true;
                const {data} = await this.$api.team.list(this.leagueId);
                this.table.data = data;
                this.table.loading = false;
            },

            /**
             * @returns {void}
             */
            async distributeFixture() {
                this.table.loading = true;
                await this.$api.match.distributeFixture(this.leagueId);
                await this.init();
            },

            /**
             * @returns {void}
             */
            async playOneWeek() {
                this.table.loading = true;
                await this.$api.match.playOneWeek(this.leagueId);
                await this.init();
            },

            /**
             * @returns {void}
             */
            async playAll() {
                this.table.loading = true;
                await this.$api.match.playAll(this.leagueId);
                await this.init();
            }
        }
    }
</script>
