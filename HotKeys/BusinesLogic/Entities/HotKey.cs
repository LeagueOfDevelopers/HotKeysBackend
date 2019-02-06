using System;
using System.Collections.Generic;
using System.Linq;

namespace HotKeysLibrary.Entities
{
    public class HotKey
    {
        public Guid Id { get; }
        public string Description { get; }
        private List<Key> _combination;
        public IEnumerable<Key> Combination => _combination.AsEnumerable();

        public HotKey(Guid id, string description, List<Key> combination)
        {
            Id = id;
            Description = description;
            _combination = combination;
        }
    }
}
