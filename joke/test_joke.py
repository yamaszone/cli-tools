import unittest
from joke import *

class test_joke(unittest.TestCase):
    def test_get_json_payload_returns_jokes_when_api_called(self):
        jokes = get_json_payload()
        self.assertIn("setup", jokes)
        self.assertIn("punchline", jokes)