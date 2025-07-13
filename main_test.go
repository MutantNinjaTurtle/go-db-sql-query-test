package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()
	clientID := 1
	client, err := selectClient(db, clientID)
	require.NoError(t, err)
	assert.Equal(t, client.ID, clientID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
	assert.NotEmpty(t, client.Login)

	// напиши тест здесь
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	clientID := -1
	client, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, client.ID)
	assert.Empty(t, client.FIO)
	assert.Empty(t, client.Login)
	assert.Empty(t, client.Birthday)
	assert.Empty(t, client.Email)
	// напиши тест здесь
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, client.ID, cl.ID)
	assert.Equal(t, client.Birthday, cl.Birthday)
	assert.Equal(t, client.Email, cl.Email)
	assert.Equal(t, client.FIO, cl.FIO)
	assert.Equal(t, client.Login, cl.Login)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	_, err = selectClient(db, id)
	require.NoError(t, err)
	err = deleteClient(db, id)
	require.NoError(t, err)
	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)

}
