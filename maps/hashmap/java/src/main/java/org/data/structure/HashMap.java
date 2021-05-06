package org.data.structure;

public class HashMap {
     class Bucket {
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
        int index = hashCode % numBuckets;
        // key.hashCode() coule be negative.
        index = index < 0 ? index * -1 : index;
        return index;
    }

    public void add(String key, int value) {
        int posic = getBucketIndex(key);
        if (this.bucket[posic] == null) {   // No collision
            Bucket newData = new Bucket(key, value);
            this.bucket[posic] = newData;
            this.numElems++;
        } else { // Collision
            Bucket existing = this.bucket[posic];
            boolean existingKey = false;

            do {
                if (existing.key.equalsIgnoreCase(key)) {
                    existingKey = true;
                    existing.value = value;
                } else {
                    existing = existing.next;
                }
            } while (existing.next != null && !existingKey);

            if (!existingKey) {
                Bucket newData = new Bucket(key, value);
                existing.next = newData;
                this.numElems++;
            }
        }

        // if load factor close to 80, increase capacity and recollocate data
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
    }

}
