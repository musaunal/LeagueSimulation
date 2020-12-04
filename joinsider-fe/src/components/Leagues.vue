<template>
    <div>
        <v-data-table :headers="table.headers" :loading="table.loading" :items="table.data">
            <template v-slot:top>
                <v-toolbar-title>Lig Tablosu</v-toolbar-title>
                <v-divider class="mx-4" inset vertical/>
                <v-spacer/>
                <v-dialog max-width="500px" v-model="dialog.show">
                    <template v-slot:activator="{on, attrs}">
                        <v-btn color="primary" dark v-bind="attrs" v-on="on">
                            Yeni lig
                        </v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <span class="headline">Lig Ekle/Düzenle</span>
                        </v-card-title>
                        <v-card-text>
                            <v-container>
                                <v-row>
                                    <v-col cols="12">
                                        <v-text-field v-model="dialog.data.name" label="Lig ismi"/>
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
                <router-link :to="`/leagues/${item.id}`">{{item.name}}</router-link>
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
        <v-overlay :value="table.loading">
            <v-progress-circular
                    indeterminate
                    size="64"
            ></v-progress-circular>
        </v-overlay>
    </div>
</template>

<script>
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
                        {text: 'Lig Ismi', value: 'name'},
                        {text: 'Oynanan Hafta', value: 'current_week'},
                        {text: 'Toplam Hafta', value: 'total_week'},
                        {text: 'Oluşturulma', value: 'created_at'},
                        {text: 'Güncelleme', value: 'updated_at'},
                        {text: 'Aksiyonlar', value: 'actions', sortable: false},
                    ],
                    data: []
                }
            }
        },
        mounted() {
            this.refresh();
        },
        methods: {
            async save() {
                this.dialog.show = false;
                this.table.loading = true;
                if(this.dialog.data.id) {
                    await this.$api.league.update(this.dialog.data.id, {name: this.dialog.data.name});
                } else {
                    await this.$api.league.create({name: this.dialog.data.name})
                }
                await this.close();
                await this.refresh();
            },
            async editItem(item) {
                this.dialog.data = item;
                console.log(item);
                this.dialog.show = true;
            },
            async deleteItem(id) {
                this.table.loading = true;
                await this.$api.league.delete(id);
                await this.refresh();
            },
            async close() {
                this.dialog.show = false;
                this.dialog.data = {}
            },
            async refresh() {
                this.table.loading = true;
                const {data} = await this.$api.league.list();
                this.table.data = data;
                this.table.loading = false;
            }
        }
    }
</script>


// [{name: "musa", current_week: "1", total_week: "32", created_at: "0",
//                                     updated_at: "4", actions: ["mdi-pencil", "mdi-delete"]  }]