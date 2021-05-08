package org.data.structure;

import java.util.Arrays;

public class HashMap {
     private class Bucket {
        String key;
        int value;
        Bucket next;

        Bucket(String key, int value) {
            this.key = key;
            this.value = value;
            this.next = null;
        }
    }

    private Bucket[] bucket;
    private int numBuckets;
    private int numElems;

    public HashMap(int numBuckets) {
        this.numBuckets = numBuckets;
        this.numElems = 0;
        this.bucket = new Bucket[numBuckets];
    }

    private int getBucketIndex(String key) {
        int hashCode = key.hashCode();
        int index = hashCode % this.numBuckets;
        // key.hashCode() coule be negative.
        index = index < 0 ? index * -1 : index;
        return index;
    }

    private void addNoCollision(String key, int value, int posic) {
        Bucket newData = new Bucket(key, value);
        this.bucket[posic] = newData;
        this.numElems++;
    }

    private void addCollision(String key, int value, int posic) {
        Bucket existing = this.bucket[posic];
        boolean existingKey = false;
        if (existing.key.equalsIgnoreCase(key)) {
            existing.value = value;
            existingKey = true;
        } else {
            while (existing.next != null && !existingKey) {
                existing = existing.next;
                if (existing.key.equalsIgnoreCase(key)) {
                    existing.value = value;
                    existingKey = true;
                }
            }
        }

        if (!existingKey) {
            Bucket newData = new Bucket(key, value);
            existing.next = newData;
            this.numElems++;
        }
    }

    public void add(String key, int value) {
        int posic = getBucketIndex(key);
        if (this.bucket[posic] == null) {   // No collision
            addNoCollision(key, value, posic);
        } else { // Collision
            addCollision(key, value, posic);
        }

        // if load factor close to 80, increase capacity and reallocate data
    }

    private boolean removeFirstElementBucket(Bucket bucket, int posic) {
        // check if there're more elements
        if (bucket.next == null) {
            this.bucket[posic] = null;
        } else {
            this.bucket[posic] = bucket.next;
        }

        this.numElems--;
        return true;
    }

    private boolean removeFromList(Bucket bucket, String key) {
        // Iterate list to find out node with key
        while (bucket.next != null) {
            if (bucket.next.key.equalsIgnoreCase(key)) {
                bucket = bucket.next.next;
                this.numElems--;
                return true;
            }

            bucket = bucket.next;
        }

        return false;
    }

    public boolean remove(String key) {
        int posic = getBucketIndex(key);
        if (this.bucket[posic] == null) {   // key does not exists
            return false;
        }

        // First element bucket
        Bucket bucket = this.bucket[posic];
        if (bucket.key.equalsIgnoreCase(key)) {
            return removeFirstElementBucket(bucket, posic);
        } else {
            return removeFromList(bucket, key);
        }
    }

    public Integer get(String key) {
        int posic = getBucketIndex(key);
        Bucket bucket = this.bucket[posic];
        if (bucket == null) {
            return null;
        }

        if (bucket.key.equalsIgnoreCase(key)) {
            return bucket.value;
        }

        while (bucket.next != null) {
            bucket = bucket.next;
            if (bucket.key.equalsIgnoreCase(key)) {
                return bucket.value;
            }
        }

        return null;
    }

    String[] getKeys() {
        if (numElems == 0) {
            return null;
        }

        int posic = 0;
        String[] keys = new String[this.numElems];
        for (int idx=0; idx<this.numBuckets; idx++) {
            if (this.bucket[idx] != null) {
                Bucket bucket = this.bucket[idx];
                keys[posic++] = bucket.key;

                while (bucket.next != null) {
                    bucket = bucket.next;
                    keys[posic++] = bucket.key;
                }
            }
        }

        return keys;
    }

    public static void main(String[] args) {
        HashMap hm = new HashMap(4);

        hm.add("key1", 1);
        hm.add("key2", 2);
        hm.add("key3", 3);
        hm.add("key4", 4);
        hm.add("key1", 14);
        hm.add("key6", 5);
        hm.add("key7", 6);
        hm.add("key8", 7);
        hm.add("key9", 8);
        hm.add("key10", 9);
        hm.add("key11", 10);
        hm.add("key12", 11);

        String[] keys = hm.getKeys();
        Arrays.stream(keys).forEach(System.out::println);

        System.out.println(hm.get("key8"));
    }

}
