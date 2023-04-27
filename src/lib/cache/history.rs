use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

pub struct History<K, V> {
    capacity: usize,
    history: Arc<Mutex<HashMap<K, Vec<V>>>>,
}

impl<K, V> History<K, V>
where
    K: Eq + std::hash::Hash + Clone,
    V: Clone,
{
    pub fn new(capacity: usize) -> Self {
        Self {
            capacity,
            history: Arc::new(Mutex::new(HashMap::new())),
        }
    }

    pub fn append(&mut self, key: K, value: V) {
        // let mut history = self.history.lock().unwrap();
        // history.get(&key)
    }

    pub fn get(self, key: &K) -> Option<V> {
        let history = self.history.lock().unwrap();
        if let Some(history) = history.get(key) {
            Some(history.get(0).unwrap().clone())
        } else {
            None
        }
    }

    pub fn back(self, key: &K, index: usize) -> Option<V> {
        let history = self.history.lock().unwrap();
        if let Some(history) = history.get(key) {
            Some(history.get(index + 1).unwrap().clone())
        } else {
            None
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn it_works() {
        let mut his = History::<String, String>::new(3);
        let v = his.get(&String::from("xx")).unwrap();
        println!("{}", v);
    }
}
