from unittest.mock import AsyncMock, patch
from fastapi.testclient import TestClient

from .main import app

client = TestClient(app)

def test_madlib_with_external_calls():
    response = client.get("/madlib")
    assert response.status_code == 200
    assert b"It was a" in response.content
    assert " day. I went downstairs to see if I could " in response.text
    assert " dinner. I asked, 'Does the stew need fresh " in response.text
    assert "?'" in response.text

@patch("app.main.fetch_word", new_callable=AsyncMock)
def test_madlib_without_external_calls(mock_fetch_word):
    mock_fetch_word.side_effect = ["sunny", "cook", "vegetables"]
    response = client.get("/madlib")
    assert response.status_code == 200
    assert "It was a sunny day. I went downstairs to see if I could cook dinner. I asked, 'Does the stew need fresh vegetables?'" in response.text

@patch("app.main.fetch_word", new_callable=AsyncMock)
def test_madlib_negative(mock_fetch_word):
    mock_fetch_word.side_effect = [ValueError("test_error"), ValueError("test_error"), ValueError("test_error")]
    response = client.get("/madlib")
    assert response.status_code == 200
    assert "test_error" in response.text