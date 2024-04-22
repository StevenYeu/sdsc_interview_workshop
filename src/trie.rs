use std::collections::HashMap;

struct WordDictionary {
    pub is_word: bool,
    pub children: HashMap<char, WordDictionary>,
}
impl WordDictionary {
    fn new() -> Self {
        return WordDictionary {
            is_word: false,
            children: HashMap::new(),
        };
    }

    fn add_word(&mut self, word: String) {
        let mut cur = self;
        for char in word.chars() {
            let letter = cur.children.get(&char);
            match letter {
                Some(_) => (),
                None => {
                    let node = WordDictionary {
                        is_word: false,
                        children: HashMap::new(),
                    };
                    cur.children.insert(char, node);
                }
            }

            cur = cur.children.get_mut(&char).unwrap()
        }
        cur.is_word = true;
    }

    fn search(&self, word: String) -> bool {
        let mut cur = self;

        for char in word.chars() {
            let letter = cur.children.get(&char);
            match letter {
                Some(a) => cur = &a,
                None => return false,
            }
        }
        cur.is_word
    }
}
