use std::collections::{HashMap, VecDeque};
use std::sync::{Arc, Mutex};

pub struct LRUCache<K, V> {
    capacity: usize,
    cache: Arc<Mutex<HashMap<K, V>>>,
    order: Arc<Mutex<VecDeque<K>>>,
}

impl<K, V> LRUCache<K, V>
where
    K: Eq + std::hash::Hash + Clone,
    V: Clone,
{
    pub fn new(capacity: usize) -> Self {
        LRUCache {
            capacity,
            cache: Arc::new(Mutex::new(HashMap::new())),
            order: Arc::new(Mutex::new(VecDeque::new())),
        }
    }

    pub fn get(&mut self, key: &K) -> Option<V> {
        let cache = self.cache.lock().unwrap();
        let mut order = self.order.lock().unwrap();
        if let Some(value) = cache.get(key).cloned() {
            // Move the key to the end of the order vector
            order.retain(|k| k != key);
            order.push_back(key.clone());
            Some(value)
        } else {
            None
        }
    }

    pub fn put(&mut self, key: K, value: V) {
        let mut cache = self.cache.lock().unwrap();
        let mut order = self.order.lock().unwrap();
        // If the key already exists, remove it from the order vector
        if cache.contains_key(&key) {
            order.retain(|k| k != &key);
        }
        // Add the key and value to the cache and the end of the order vector
        cache.insert(key.clone(), value);
        order.push_back(key.clone());
        // If the cache size exceeds the capacity, remove the least recently used key
        while cache.len() > self.capacity {
            if let Some(key) = order.pop_front() {
                cache.remove(&key);
            }
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn it_works() {
        let mut lru = LRUCache::new(3);
        lru.put("1", "x");
        lru.put("2", "x2");
        lru.put("3", "x3");
        lru.put("4", "x4");
        let k = lru.get(&"1").unwrap();
        println!("{}", k);
    }
}
