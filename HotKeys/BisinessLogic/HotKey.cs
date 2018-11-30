using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    public class HotKey
    {
        public Guid Id { get; }
        public string Description { get; }
        public Dictionary<int, string> Combination;
        public HotKey(Guid id, string description, Dictionary<int, string> combination)
        {
            Id = id;
            Description = description;
            Combination = combination;
        }

    }
}
