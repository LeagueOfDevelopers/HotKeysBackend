using HotKeysLibrary.Entities;
using EnsureThat;
using System;
using System.Collections.Generic;
using System.Text;

namespace HotKeysLibrary.Infrastructure
{
    class InMemoryUserRepository : IRepository<User>
    {
        private readonly Dictionary<Guid, User> _users;

        public InMemoryUserRepository()
        {
            _users = new Dictionary<Guid, User>();
        }

        public User Get(Guid id)
        {
            return Ensure.Any.IsNotNull(_users[id], nameof(Get),
                opt => opt.WithException(new Exception())); //заменить Exception
        }

        public void Add(Guid id, User user)
        {
            if (user == null)
                throw new ArgumentNullException();
            _users.Add(id, user);
        }
    }
}
