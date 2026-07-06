package main

import (
	"LBD/database"
	"LBD/schemas"
	"LBD/seed"
	"LBD/services"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&schemas.Artista{},
		&schemas.Usuario{},
		&schemas.Musica{},
		&schemas.Playlist{},
		&schemas.MusicaPlaylist{},
	)
	if err != nil {
		panic(err)
	}

	if err := applyManualConstraints(db); err != nil {
		panic(err)
	}

	modo := "demo"
	if len(os.Args) > 1 {
		modo = os.Args[1]
	}

	sufixo := time.Now().Format("20060102150405")
	args := os.Args[2:]

	switch modo {
	case "help":
		printUso()
		return
	case "demo":
		if err := executarDemoCompleta(db, sufixo); err != nil {
			panic(err)
		}
		return
	case "smoke":
		if err := executarSmokeTest(db, sufixo); err != nil {
			panic(err)
		}
		return
	case "seed":
		seed.Seed(db)
		fmt.Println("Seed concluído")
		return
	case "createArtista":
		if err := executarCreateArtista(sufixo); err != nil {
			panic(err)
		}
	case "resetDb":
		if err := resetDB(db); err != nil {
			panic(err)
		}
		fmt.Println("resetDb concluído")
		return
	case "getArtista":
		if err := executarGetArtista(args); err != nil {
			panic(err)
		}
	case "getAllArtista":
		if err := executarGetAllArtista(); err != nil {
			panic(err)
		}
	case "updateArtista":
		if err := executarUpdateArtista(args, sufixo); err != nil {
			panic(err)
		}
	case "deleteArtista":
		if err := executarDeleteArtista(args); err != nil {
			panic(err)
		}
	case "createMusica":
		if err := executarCreateMusica(args, sufixo); err != nil {
			panic(err)
		}
	case "getMusica":
		if err := executarGetMusica(args); err != nil {
			panic(err)
		}
	case "getAllMusica":
		if err := executarGetAllMusica(); err != nil {
			panic(err)
		}
	case "updateMusica":
		if err := executarUpdateMusica(args, sufixo); err != nil {
			panic(err)
		}
	case "deleteMusica":
		if err := executarDeleteMusica(args); err != nil {
			panic(err)
		}
	case "createPlaylist":
		if err := executarCreatePlaylist(args, sufixo); err != nil {
			panic(err)
		}
	case "addMusicaToPlaylist":
		if err := executarAddMusicaToPlaylist(args); err != nil {
			panic(err)
		}
	case "removeMusicaFromPlaylist":
		if err := executarRemoveMusicaFromPlaylist(args); err != nil {
			panic(err)
		}
	case "transferirMusicaEntrePlaylists":
		if err := executarTransferirMusicaEntrePlaylists(args); err != nil {
			panic(err)
		}
	case "getAllPlaylist":
		if err := executarGetAllPlaylist(); err != nil {
			panic(err)
		}
	case "listarRockDoPablo":
		if err := executarListarMusicasRockDoPablo(); err != nil {
			panic(err)
		}
	case "listarDonoBohemianRhapsody":
		if err := executarListarDonoDaMusica("Bohemian Rhapsody"); err != nil {
			panic(err)
		}
	case "getMusicasLedZeppelinMaioresQueQueen":
		if err := GetMusicasLedZeppelinMaioresQueQueen(); err != nil {
			panic(err)
		}
	case "getRankingPopularidadeArtista":
		if err := GetRankingPopularidadeArtista(); err != nil {
			panic(err)
		}
	case "getMusicasByUsuarioAndArtista":
		if err := executarGetMusicasByUsuarioAndArtista(args); err != nil {
			panic(err)
		}
	case "getQuantidadeMusicasPorPlaylist":
		if err := executarGetQuantidadeMusicasPorPlaylist(); err != nil {
			panic(err)
		}
	case "getArtistasSemMusicasEmPlaylist":
		if err := executarGetArtistasSemMusicasEmPlaylist(); err != nil {
			panic(err)
		}
	case "getMusicaComArtista":
		if err := executarGetMusicaComArtista(args); err != nil {
			panic(err)
		}
	case "getTempoTotalPorPlaylist":
		if err := executarGetTempoTotalPorPlaylist(); err != nil {
			panic(err)
		}
	case "getMusicasMaisCurtasQueMediaDoArtista":
		if err := executarGetMusicasMaisCurtasQueMediaDoArtista(); err != nil {
			panic(err)
		}
	case "getAllUsuario":
		if err := executarGetAllUsuario(); err != nil {
			panic(err)
		}
	case "artista":
		if err := executarTesteArtista(sufixo); err != nil {
			panic(err)
		}
	case "musica":
		if err := executarTesteMusica(sufixo); err != nil {
			panic(err)
		}
	case "playlist":
		if err := executarTestePlaylist(sufixo); err != nil {
			panic(err)
		}
	case "imprimirArtistas":
		artistas, err := services.GetAllArtista()
		if err != nil {
			panic(err)
		}
		imprimirArtistas(artistas)
	case "imprimirMusicas":
		musicas, err := services.GetAllMusica()
		if err != nil {
			panic(err)
		}
		imprimirMusicas(musicas)
	case "all":
		if err := executarDemoCompleta(db, sufixo); err != nil {
			panic(err)
		}
	default:
		printUso()
	}
}

func printUso() {
	fmt.Println("Uso:")
	fmt.Println("  go run .")
	fmt.Println("  go run . demo")
	fmt.Println("  go run . smoke")
	fmt.Println("  go run . resetDb")
	fmt.Println("  go run . seed")
	fmt.Println("  go run . createArtista")
	fmt.Println("  go run . createMusica <artistaID>")
	fmt.Println("  go run . createPlaylist <usuarioID>")
	fmt.Println("  go run . addMusicaToPlaylist <musicaID> <playlistID> <usuarioID>")
	fmt.Println("  go run . removeMusicaFromPlaylist <musicaID> <playlistID> <usuarioID>")
	fmt.Println("  go run . transferirMusicaEntrePlaylists <musicaID> <playlistOrigemID> <playlistDestinoID> <usuarioID>")
	fmt.Println("  go run . listarRockDoPablo")
	fmt.Println("  go run . listarDonoBohemianRhapsody")
	fmt.Println("  go run . getMusicasLedZeppelinMaioresQueQueen")
	fmt.Println("  go run . getRankingPopularidadeArtista")
	fmt.Println("  go run . getMusicasByUsuarioAndArtista <username> <nomeArtista>")
	fmt.Println("  go run . getQuantidadeMusicasPorPlaylist")
	fmt.Println("  go run . getArtistasSemMusicasEmPlaylist")
	fmt.Println("  go run . getMusicaComArtista <musicaID>")
	fmt.Println("  go run . getTempoTotalPorPlaylist")
	fmt.Println("  go run . getMusicasMaisCurtasQueMediaDoArtista")
	fmt.Println("  go run . help")
}

func printResultado(nome string, err error) {
	if err != nil {
		fmt.Printf("[ERRO] %s: %v\n", nome, err)
		return
	}

	fmt.Printf("[OK] %s\n", nome)
}

func imprimirArtistas(artistas []schemas.Artista) {
	if len(artistas) == 0 {
		fmt.Println("Nenhum artista encontrado.")
		return
	}

	fmt.Println("Lista de artistas:")
	for _, artista := range artistas {
		fmt.Printf("- ID=%d | Nome=%s | Nacionalidade=%s\n", artista.ID, artista.Nome, artista.Nacionalidade.String)
	}
}

func imprimirMusicas(musicas []schemas.Musica) {
	if len(musicas) == 0 {
		fmt.Println("Nenhuma musica encontrada.")
		return
	}

	fmt.Println("Lista de musicas:")
	for _, musica := range musicas {
		fmt.Printf("- ID=%d | Titulo=%s | Duracao=%d | ArtistaID=%d\n", musica.ID, musica.Titulo, musica.Duracao_segundos, musica.Artista_id)
	}
}

func imprimirPlaylists(playlists []schemas.Playlist) {
	if len(playlists) == 0 {
		fmt.Println("Nenhuma playlist encontrada.")
		return
	}

	fmt.Println("Lista de playlists:")
	for _, p := range playlists {
		fmt.Printf("- PlaylistId=%d | UsuarioId=%d | Nome=%s | DataCriacao=%s\n", p.PlaylistId, p.UsuarioId, p.Nome, p.DataCriacao.String())
	}
}

func imprimirUsuarios(usuarios []schemas.Usuario) {
	if len(usuarios) == 0 {
		fmt.Println("Nenhum usuario encontrado.")
		return
	}

	fmt.Println("Lista de usuarios:")
	for _, u := range usuarios {
		fmt.Printf("- ID=%d | Username=%s | Email=%s\n", u.ID, u.Username, u.Email)
	}
}

func imprimirMusicasDaPlaylist(nomePlaylist string, musicas []services.MusicaNaPlaylist) {
	if len(musicas) == 0 {
		fmt.Printf("Nenhuma musica encontrada na playlist %q.\n", nomePlaylist)
		return
	}

	fmt.Printf("Musicas na playlist %q:\n", nomePlaylist)
	for _, musica := range musicas {
		fmt.Printf("- Ordem=%d | Titulo=%s\n", musica.OrdemNaPlaylist, musica.Titulo)
	}
}

func executarListarMusicasRockDoPablo() error {
	musicas, err := services.GetMusicasDaPlaylist("Rock do Pablo")
	printResultado("GetMusicasDaPlaylist", err)
	if err != nil {
		return err
	}

	imprimirMusicasDaPlaylist("Rock do Pablo", musicas)
	return nil
}

func executarListarDonoDaMusica(tituloMusica string) error {
	usuarios, err := services.GetUsuariosDonoDaMusica(tituloMusica)
	printResultado("GetUsuariosDonoDaMusica", err)
	if err != nil {
		return err
	}

	if len(usuarios) == 0 {
		fmt.Printf("Nenhum usuario encontrado para a musica %q.\n", tituloMusica)
		return nil
	}

	fmt.Printf("Usuarios donos da playlist que contem %q:\n", tituloMusica)
	for _, username := range usuarios {
		fmt.Printf("- Username=%s\n", username)
	}

	return nil
}

func GetMusicasLedZeppelinMaioresQueQueen() error {
	musicas, err := services.GetMusicasLedZeppelinMaioresQueMaiorMusicaQueen()
	printResultado("GetMusicasLedZeppelinMaioresQueMaiorMusicaQueen", err)
	if err != nil {
		return err
	}

	if len(musicas) == 0 {
		fmt.Println("Nenhuma musica do Led Zeppelin encontrada com duracao maior que a musica mais longa do Queen.")
		return nil
	}

	fmt.Println("Musicas do Led Zeppelin com duracao maior que a musica mais longa do Queen:")
	for _, musica := range musicas {
		fmt.Printf("- ID=%d | Titulo=%s | Duracao=%d | ArtistaID=%d\n", musica.ID, musica.Titulo, musica.Duracao_segundos, musica.Artista_id)
	}

	return nil
}

func GetRankingPopularidadeArtista() error {
	ranking, err := services.GetRankingPopularidadeArtista()
	printResultado("GetRankingPopularidadeArtista", err)
	if err != nil {
		return err
	}

	if len(ranking) == 0 {
		fmt.Println("Nenhum artista encontrado.")
		return nil
	}

	fmt.Println("Ranking de popularidade dos artistas:")
	for _, item := range ranking {
		fmt.Printf(
			"%dº - ArtistaID=%d | Nome=%s | Playlists=%d\n",
			item.Posicao,
			item.ArtistaID,
			item.Nome,
			item.QuantidadePlaylists,
		)
	}

	return nil
}

func executarGetMusicasByUsuarioAndArtista(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("uso: getMusicasByUsuarioAndArtista <username> <nomeArtista>")
	}

	musicas, err := services.GetMusicasByUsuarioAndArtista(args[0], args[1])
	printResultado("GetMusicasByUsuarioAndArtista", err)
	if err != nil {
		return err
	}

	imprimirMusicas(musicas)
	return nil
}

func executarGetQuantidadeMusicasPorPlaylist() error {
	resultado, err := services.GetQuantidadeMusicasPorPlaylist()
	printResultado("GetQuantidadeMusicasPorPlaylist", err)
	if err != nil {
		return err
	}

	if len(resultado) == 0 {
		fmt.Println("Nenhuma playlist encontrada.")
		return nil
	}

	fmt.Println("Quantidade de musicas por playlist:")
	for _, item := range resultado {
		fmt.Printf("- Playlist=%s | QuantidadeMusicas=%d\n", item.Nome, item.QuantidadeMusicas)
	}
	return nil
}

func executarGetArtistasSemMusicasEmPlaylist() error {
	artistas, err := services.GetArtistasSemMusicasEmPlaylist()
	printResultado("GetArtistasSemMusicasEmPlaylist", err)
	if err != nil {
		return err
	}

	imprimirArtistas(artistas)
	return nil
}

func executarGetMusicaComArtista(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}

	musica, err := services.GetMusicaComArtista(musicaID)
	printResultado("GetMusicaComArtista", err)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Musica: ID=%d Titulo=%s Duracao=%d | Artista: ID=%d Nome=%s\n",
		musica.ID, musica.Titulo, musica.Duracao_segundos, musica.Artista.ID, musica.Artista.Nome,
	)
	return nil
}

func executarGetTempoTotalPorPlaylist() error {
	resultado, err := services.GetTempoTotalPorPlaylist()
	printResultado("GetTempoTotalPorPlaylist", err)
	if err != nil {
		return err
	}

	if len(resultado) == 0 {
		fmt.Println("Nenhuma playlist encontrada.")
		return nil
	}

	fmt.Println("Tempo total de reproducao por playlist:")
	for _, item := range resultado {
		fmt.Printf("- Playlist=%s | Dono=%s | TempoTotalSegundos=%d\n", item.NomePlaylist, item.UsernameDono, item.TempoTotalSegundos)
	}
	return nil
}

func executarGetMusicasMaisCurtasQueMediaDoArtista() error {
	musicas, err := services.GetMusicasMaisCurtasQueMediaDoArtista()
	printResultado("GetMusicasMaisCurtasQueMediaDoArtista", err)
	if err != nil {
		return err
	}

	imprimirMusicas(musicas)
	return nil
}

func resetDB(db *gorm.DB) error {
	fmt.Println("=== RESET DB: Dropping tables and reapplying migrations ===")

	for _, query := range []string{
		`DROP TABLE IF EXISTS "MUSICA_PLAYLIST" CASCADE`,
		`DROP TABLE IF EXISTS "PLAYLIST" CASCADE`,
		`DROP TABLE IF EXISTS "MUSICA" CASCADE`,
		`DROP TABLE IF EXISTS "ARTISTA" CASCADE`,
		`DROP TABLE IF EXISTS "USUARIO" CASCADE`,
		`DROP TABLE IF EXISTS "musica_playlists" CASCADE`,
		`DROP TABLE IF EXISTS "playlists" CASCADE`,
		`DROP TABLE IF EXISTS "musicas" CASCADE`,
		`DROP TABLE IF EXISTS "artistas" CASCADE`,
		`DROP TABLE IF EXISTS "usuarios" CASCADE`,
	} {
		if err := db.Exec(query).Error; err != nil {
			return err
		}
	}

	if err := db.AutoMigrate(
		&schemas.Artista{},
		&schemas.Usuario{},
		&schemas.Musica{},
		&schemas.Playlist{},
		&schemas.MusicaPlaylist{},
	); err != nil {
		return err
	}

	return applyManualConstraints(db)
}

func executarDemoCompleta(db *gorm.DB, sufixo string) error {
	if err := resetDB(db); err != nil {
		return err
	}

	seed.Seed(db)
	fmt.Println("Seed concluído")

	if err := executarSmokeTest(db, sufixo); err != nil {
		return err
	}

	return nil
}

func executarSmokeTest(db *gorm.DB, sufixo string) error {
	fmt.Println("=== SMOKE TEST ===")

	if err := executarTesteArtista(sufixo); err != nil {
		return err
	}

	if err := executarTesteMusica(sufixo); err != nil {
		return err
	}

	if err := executarTestePlaylist(sufixo); err != nil {
		return err
	}

	return executarTesteDelete(db, sufixo)
}

func applyManualConstraints(db *gorm.DB) error {
	queries := []string{
		`ALTER TABLE "MUSICA_PLAYLIST" DROP CONSTRAINT IF EXISTS "fk_playlist_composta"`,
		`ALTER TABLE "MUSICA_PLAYLIST" ADD CONSTRAINT "fk_playlist_composta" FOREIGN KEY (playlist_id, usuario_id) REFERENCES "PLAYLIST" (playlist_id, usuario_id) ON DELETE CASCADE`,
	}

	for _, query := range queries {
		if err := db.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}

func executarCreateArtista(sufixo string) error {
	fmt.Println("=== CREATE ARTISTA ===")

	novoArtista := schemas.Artista{
		Nome: "Artista Criado " + sufixo,
		Nacionalidade: sql.NullString{
			String: "Brasileira",
			Valid:  true,
		},
	}

	err := services.CreateArtista(novoArtista)
	printResultado("CreateArtista", err)
	if err != nil {
		return err
	}

	fmt.Printf("Artista criado: Nome=%s\n", novoArtista.Nome)
	return nil
}

func executarGetArtista(args []string) error {
	artistaID, err := parseUintArg(args, 0, "artistaID")
	if err != nil {
		return err
	}

	artista, err := services.GetArtista(artistaID)
	printResultado("GetArtista", err)
	if err != nil {
		return err
	}

	fmt.Printf("Artista encontrado: ID=%d Nome=%s\n", artista.ID, artista.Nome)
	return nil
}

func executarGetAllArtista() error {
	artistas, err := services.GetAllArtista()
	printResultado("GetAllArtista", err)
	if err != nil {
		return err
	}

	imprimirArtistas(artistas)
	return nil
}

func executarUpdateArtista(args []string, sufixo string) error {
	artistaID, err := parseUintArg(args, 0, "artistaID")
	if err != nil {
		return err
	}

	artistaAtualizado, err := services.UpdateArtista(schemas.Artista{
		ID:   artistaID,
		Nome: "Artista Atualizado " + sufixo,
		Nacionalidade: sql.NullString{
			String: "Portuguesa",
			Valid:  true,
		},
	})
	printResultado("UpdateArtista", err)
	if err != nil {
		return err
	}

	fmt.Printf("Artista atualizado: ID=%d Nome=%s\n", artistaAtualizado.ID, artistaAtualizado.Nome)
	return nil
}

func executarDeleteArtista(args []string) error {
	artistaID, err := parseUintArg(args, 0, "artistaID")
	if err != nil {
		return err
	}

	err = services.DeleteArtista(artistaID)
	printResultado("DeleteArtista", err)
	return err
}

func executarCreateMusica(args []string, sufixo string) error {
	artistaID, err := parseUintArg(args, 0, "artistaID")
	if err != nil {
		return err
	}

	tituloMusica := "Musica Criada " + sufixo
	err = services.CreateMusica(schemas.Musica{
		Titulo:           tituloMusica,
		Duracao_segundos: 210,
		Artista_id:       artistaID,
	})
	printResultado("CreateMusica", err)
	if err != nil {
		return err
	}

	fmt.Printf("Musica criada: Titulo=%s ArtistaID=%d\n", tituloMusica, artistaID)
	return nil
}

func executarGetMusica(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}

	musica, err := services.GetMusica(musicaID)
	printResultado("GetMusica", err)
	if err != nil {
		return err
	}

	fmt.Printf("Musica encontrada: ID=%d Titulo=%s\n", musica.ID, musica.Titulo)
	return nil
}

func executarGetAllMusica() error {
	musicas, err := services.GetAllMusica()
	printResultado("GetAllMusica", err)
	if err != nil {
		return err
	}

	imprimirMusicas(musicas)
	return nil
}

func executarGetAllPlaylist() error {
	playlists, err := services.GetAllPlaylist()
	printResultado("GetAllPlaylist", err)
	if err != nil {
		return err
	}

	imprimirPlaylists(playlists)
	return nil
}

func executarGetAllUsuario() error {
	usuarios, err := services.GetAllUsuario()
	printResultado("GetAllUsuario", err)
	if err != nil {
		return err
	}

	imprimirUsuarios(usuarios)
	return nil
}

func executarUpdateMusica(args []string, sufixo string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}
	artistaID, err := parseUintArg(args, 1, "artistaID")
	if err != nil {
		return err
	}

	musicaAtualizada, err := services.UpdateMusica(schemas.Musica{
		ID:               musicaID,
		Titulo:           "Musica Atualizada " + sufixo,
		Duracao_segundos: 240,
		Artista_id:       artistaID,
	})
	printResultado("UpdateMusica", err)
	if err != nil {
		return err
	}

	fmt.Printf("Musica atualizada: ID=%d Titulo=%s\n", musicaAtualizada.ID, musicaAtualizada.Titulo)
	return nil
}

func executarDeleteMusica(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}

	err = services.DeleteMusica(musicaID)
	printResultado("DeleteMusica", err)
	return err
}

func executarCreatePlaylist(args []string, sufixo string) error {
	usuarioID, err := parseUintArg(args, 0, "usuarioID")
	if err != nil {
		return err
	}

	playlist, err := services.CreatePlaylist(schemas.Playlist{
		UsuarioId: usuarioID,
		Nome:      "Playlist Criada " + sufixo,
	})
	printResultado("CreatePlaylist", err)
	if err != nil {
		return err
	}

	fmt.Printf("Playlist criada: PlaylistId=%d UsuarioId=%d Nome=%s\n", playlist.PlaylistId, playlist.UsuarioId, playlist.Nome)
	return nil
}

func executarAddMusicaToPlaylist(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}
	playlistID, err := parseUintArg(args, 1, "playlistID")
	if err != nil {
		return err
	}
	usuarioID, err := parseUintArg(args, 2, "usuarioID")
	if err != nil {
		return err
	}

	err = services.AddMusicaToPlaylist(musicaID, playlistID, usuarioID)
	printResultado("AddMusicaToPlaylist", err)
	return err
}

func executarRemoveMusicaFromPlaylist(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}
	playlistID, err := parseUintArg(args, 1, "playlistID")
	if err != nil {
		return err
	}
	usuarioID, err := parseUintArg(args, 2, "usuarioID")
	if err != nil {
		return err
	}

	err = services.RemoveMusicaFromPlaylist(musicaID, playlistID, usuarioID)
	printResultado("RemoveMusicaFromPlaylist", err)
	return err
}

func executarTransferirMusicaEntrePlaylists(args []string) error {
	musicaID, err := parseUintArg(args, 0, "musicaID")
	if err != nil {
		return err
	}
	playlistOrigemID, err := parseUintArg(args, 1, "playlistOrigemID")
	if err != nil {
		return err
	}
	playlistDestinoID, err := parseUintArg(args, 2, "playlistDestinoID")
	if err != nil {
		return err
	}
	usuarioID, err := parseUintArg(args, 3, "usuarioID")
	if err != nil {
		return err
	}

	err = services.TransferirMusicaEntrePlaylists(musicaID, playlistOrigemID, playlistDestinoID, usuarioID)
	printResultado("TransferirMusicaEntrePlaylists", err)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Musica transferida: MusicaId=%d | OrigemPlaylistId=%d | DestinoPlaylistId=%d | UsuarioId=%d\n",
		musicaID,
		playlistOrigemID,
		playlistDestinoID,
		usuarioID,
	)
	return nil
}

func parseUintArg(args []string, index int, nome string) (uint, error) {
	if len(args) <= index {
		return 0, fmt.Errorf("faltando argumento %s", nome)
	}

	valor, err := strconv.ParseUint(args[index], 10, 0)
	if err != nil {
		return 0, fmt.Errorf("argumento %s invalido: %w", nome, err)
	}

	return uint(valor), nil
}

func executarTodosOsTestes(db *gorm.DB, sufixo string) error {
	if err := executarTesteArtista(sufixo); err != nil {
		return err
	}

	if err := executarTesteMusica(sufixo); err != nil {
		return err
	}

	if err := executarTestePlaylist(sufixo); err != nil {
		return err
	}

	return executarTesteDelete(db, sufixo)
}

func executarTesteArtista(sufixo string) error {
	fmt.Println("=== TESTES DO SERVICE DE ARTISTA ===")

	nomeArtista := "Artista Teste " + sufixo
	err := services.CreateArtista(schemas.Artista{
		Nome: nomeArtista,
		Nacionalidade: sql.NullString{
			String: "Brasileira",
			Valid:  true,
		},
	})
	printResultado("CreateArtista", err)
	if err != nil {
		return err
	}

	artistas, err := services.GetAllArtista()
	printResultado("GetAllArtista", err)
	if err != nil {
		return err
	}
	imprimirArtistas(artistas)

	var artistaID uint
	for _, artista := range artistas {
		if artista.Nome == nomeArtista {
			artistaID = artista.ID
			break
		}
	}

	artista, err := services.GetArtista(artistaID)
	printResultado("GetArtista", err)
	if err != nil {
		return err
	}
	fmt.Printf("Artista encontrado: ID=%d Nome=%s\n", artista.ID, artista.Nome)

	artistaAtualizado, err := services.UpdateArtista(schemas.Artista{
		ID:   artistaID,
		Nome: nomeArtista + " Atualizado",
		Nacionalidade: sql.NullString{
			String: "Portuguesa",
			Valid:  true,
		},
	})
	printResultado("UpdateArtista", err)
	if err != nil {
		return err
	}
	fmt.Printf("Artista atualizado: ID=%d Nome=%s\n", artistaAtualizado.ID, artistaAtualizado.Nome)

	//err = services.DeleteArtista(artistaID)
	printResultado("DeleteArtista", err)
	return err
}

func executarTesteMusica(sufixo string) error {
	fmt.Println("\n=== TESTES DO SERVICE DE MUSICA ===")

	artistaID, err := criarArtistaBase(sufixo)
	if err != nil {
		return err
	}

	tituloMusica := "Musica Teste " + sufixo
	err = services.CreateMusica(schemas.Musica{
		Titulo:           tituloMusica,
		Duracao_segundos: 210,
		Artista_id:       artistaID,
	})
	printResultado("CreateMusica", err)
	if err != nil {
		_ = services.DeleteArtista(artistaID)
		return err
	}

	musicas, err := services.GetAllMusica()
	printResultado("GetAllMusica", err)
	if err != nil {
		_ = services.DeleteMusica(0)
		_ = services.DeleteArtista(artistaID)
		return err
	}
	//imprimirMusicas(musicas)

	var musicaID uint
	for _, musica := range musicas {
		if musica.Titulo == tituloMusica {
			musicaID = musica.ID
			break
		}
	}

	musica, err := services.GetMusica(musicaID)
	printResultado("GetMusica", err)
	if err != nil {
		_ = services.DeleteArtista(artistaID)
		return err
	}
	fmt.Printf("Musica encontrada: ID=%d Titulo=%s\n", musica.ID, musica.Titulo)

	musicaAtualizada, err := services.UpdateMusica(schemas.Musica{
		ID:               musicaID,
		Titulo:           tituloMusica + " Atualizada",
		Duracao_segundos: 240,
		Artista_id:       artistaID,
	})
	printResultado("UpdateMusica", err)
	if err != nil {
		_ = services.DeleteMusica(musicaID)
		_ = services.DeleteArtista(artistaID)
		return err
	}
	fmt.Printf("Musica atualizada: ID=%d Titulo=%s\n", musicaAtualizada.ID, musicaAtualizada.Titulo)

	err = services.DeleteMusica(musicaID)
	printResultado("DeleteMusica", err)
	if err != nil {
		_ = services.DeleteArtista(artistaID)
		return err
	}

	err = services.DeleteArtista(artistaID)
	printResultado("DeleteArtista", err)
	return err
}

func executarTestePlaylist(sufixo string) error {
	fmt.Println("\n=== TESTES DO SERVICE DE PLAYLIST ===")

	usuario, err := criarUsuarioTeste(sufixo)
	if err != nil {
		return err
	}
	defer func() {
		db, dbErr := database.ConnectDB()
		if dbErr == nil {
			db.Delete(&schemas.Usuario{}, usuario.ID)
		}
	}()

	artistaID, err := criarArtistaBase(sufixo)
	if err != nil {
		return err
	}
	defer func() {
		_ = services.DeleteArtista(artistaID)
	}()

	musicaID, err := criarMusicaBase(sufixo, artistaID)
	if err != nil {
		return err
	}
	defer func() {
		_ = services.DeleteMusica(musicaID)
	}()

	playlist, err := services.CreatePlaylist(schemas.Playlist{
		UsuarioId: usuario.ID,
		Nome:      "Playlist Teste " + sufixo,
	})
	printResultado("CreatePlaylist", err)
	if err != nil {
		return err
	}
	fmt.Printf("Playlist criada: PlaylistId=%d UsuarioId=%d Nome=%s\n", playlist.PlaylistId, playlist.UsuarioId, playlist.Nome)

	err = services.AddMusicaToPlaylist(musicaID, playlist.PlaylistId, playlist.UsuarioId)
	printResultado("AddMusicaToPlaylist", err)
	if err != nil {
		return err
	}

	err = services.RemoveMusicaFromPlaylist(musicaID, playlist.PlaylistId, playlist.UsuarioId)
	printResultado("RemoveMusicaFromPlaylist", err)
	if err != nil {
		return err
	}

	err = removerPlaylist(playlist)
	printResultado("DeletePlaylist", err)
	return err
}

func executarTesteDelete(db *gorm.DB, sufixo string) error {
	fmt.Println("\n=== TESTES DE DELETE ===")

	artistaID, err := criarArtistaBase(sufixo)
	if err != nil {
		return err
	}

	musicaID, err := criarMusicaBase(sufixo, artistaID)
	if err != nil {
		_ = services.DeleteArtista(artistaID)
		return err
	}

	err = services.DeleteMusica(musicaID)
	printResultado("DeleteMusica", err)
	if err != nil {
		_ = services.DeleteArtista(artistaID)
		return err
	}

	err = services.DeleteArtista(artistaID)
	printResultado("DeleteArtista", err)
	return err
}

func criarUsuarioTeste(sufixo string) (schemas.Usuario, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return schemas.Usuario{}, err
	}

	usuario := schemas.Usuario{
		Username: "Usuario Teste " + sufixo,
		Email:    "usuario.teste." + sufixo + "@email.com",
	}

	if err := db.Create(&usuario).Error; err != nil {
		return schemas.Usuario{}, err
	}

	return usuario, nil
}

func criarArtistaBase(sufixo string) (uint, error) {
	nomeArtista := "Artista Teste " + sufixo
	err := services.CreateArtista(schemas.Artista{
		Nome: nomeArtista,
		Nacionalidade: sql.NullString{
			String: "Brasileira",
			Valid:  true,
		},
	})
	printResultado("CreateArtista", err)
	if err != nil {
		return 0, err
	}

	artistas, err := services.GetAllArtista()
	printResultado("GetAllArtista", err)
	if err != nil {
		return 0, err
	}
	imprimirArtistas(artistas)

	for _, artista := range artistas {
		if artista.Nome == nomeArtista {
			return artista.ID, nil
		}
	}

	return 0, fmt.Errorf("artista teste nao encontrado")
}

func criarMusicaBase(sufixo string, artistaID uint) (uint, error) {
	tituloMusica := "Musica Teste " + sufixo
	err := services.CreateMusica(schemas.Musica{
		Titulo:           tituloMusica,
		Duracao_segundos: 210,
		Artista_id:       artistaID,
	})
	printResultado("CreateMusica", err)
	if err != nil {
		return 0, err
	}

	musicas, err := services.GetAllMusica()
	printResultado("GetAllMusica", err)
	if err != nil {
		return 0, err
	}
	imprimirMusicas(musicas)

	for _, musica := range musicas {
		if musica.Titulo == tituloMusica {
			return musica.ID, nil
		}
	}

	return 0, fmt.Errorf("musica teste nao encontrada")
}

func removerPlaylist(playlist schemas.Playlist) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	err = db.Where("playlist_id = ? AND usuario_id = ?", playlist.PlaylistId, playlist.UsuarioId).Delete(&schemas.MusicaPlaylist{}).Error
	if err != nil {
		return err
	}

	return db.Where("playlist_id = ? AND usuario_id = ?", playlist.PlaylistId, playlist.UsuarioId).Delete(&schemas.Playlist{}).Error
}
